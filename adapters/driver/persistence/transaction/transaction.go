package transaction

import (
	"database/sql"
	"errors"
	"fmt"
	"math"
	"time"

	dbErrors "github.com/NicklasWallgren/go-template/adapters/driver/persistence/errors"
	"gorm.io/gorm"
)

// Source https://github.com/mightyguava/autotx with added changes
// nolint:wsl

// DefaultMaxRetries configures the default number of max retries attempted by TransactWithRetry.
var DefaultMaxRetries = 1

// DefaultIsRetryable configures the default function for determining whether the error returned from the operation is
// retryable. By default, applicable errors are retryable. A RollbackError is never retryable..
var DefaultIsRetryable = retryIfApplicable

// Transact executes the operation inside a transaction, committing the transaction on completion. If the operation
// returns an error or panic, the transaction will be rolled back, returning the original error or propagating the
// original panic. If the rollback caused by an error also receives an error, a RollbackError will be returned. If the
// rollback caused by a panic returns an error, the error message and original panic merged and propagated as a new
// panic.
func Transact(db *gorm.DB, operation func(tx *gorm.DB) error) (err error) {
	return TransactWithOptions(db, nil, operation)
}

// TransactWithOptions executes the operation inside a transaction, committing the transaction on completion. If the
// operation returns an error or panic, the transaction will be rolled back, returning the original error or propagating
// the original panic. If the rollback caused by an error also receives an error, a RollbackError will be returned. If the
// rollback caused by a panic returns an error, the error message and original panic merged and propagated as a new
// panic.
//
// The provided TxOptions is optional and may be nil if defaults should be used. If a non-default isolation level is
// used that the driver doesn't support, an error will be returned.
func TransactWithOptions(db *gorm.DB, txOpts *sql.TxOptions, operation func(tx *gorm.DB) error) (err error) {
	tx := db.Begin(txOpts)

	defer func() {
		if p := recover(); p != nil {
			if tx.Rollback(); tx.Error != nil {
				p = fmt.Errorf("panic in transaction, AND rollback failed with error: %v, original panic: %v", tx.Error, p)
			}

			panic(p)
		}
		if err != nil {
			if tx.Rollback(); tx.Error != nil {
				err = &RollbackError{
					OriginalErr: err,
					Err:         tx.Error,
				}
			}
			return
		}
		err = tx.Commit().Error
	}()

	err = operation(tx)
	return
}

// TransactWithRetry runs the operation using Transact, performing retries according to RetryOptions. If all retries
// fail, the error from the last attempt will be returned. If a rollback fails, no further attempts will be made and the
// RollbackError will be returned.
//
// Since the transaction operation may be executed multiple times, it is important that any mutations it applies
// to application state (outside the database) be idempotent.
func TransactWithRetry(db *gorm.DB, retry RetryOptions, operation func(tx *gorm.DB) error) error {
	// TODO, what happens if a transaction already has been initialized

	return TransactWithRetryAndOptions(db, nil, retry, operation)
}

// TransactWithDefaultRetry runs the operation using Transact, performing retries according to RetryOptions. If all retries
// fail, the error from the last attempt will be returned. If a rollback fails, no further attempts will be made and the
// RollbackError will be returned.
//
// Since the transaction operation may be executed multiple times, it is important that any mutations it applies
// to application state (outside the database) be idempotent.
func TransactWithDefaultRetry(db *gorm.DB, operation func(tx *gorm.DB) error) error {
	// TODO, default retry for locking and duplicate keys

	return TransactWithRetryAndOptions(db, nil, RetryOptions{}, operation)
}

// TransactWithRetryAndOptions runs the operation using Transact, performing retries according to RetryOptions. If all
// retries fail, the error from the last attempt will be returned. If a rollback fails, no further attempts will be made
// and the RollbackError will be returned.
//
// Since the transaction operation may be executed multiple times, it is important that any mutations it applies to
// application state (outside the database) be idempotent.
//
// The provided TxOptions is optional and may be nil if defaults should be used. If a non-default isolation level is
// used that the driver doesn't support, an error will be returned.
func TransactWithRetryAndOptions(db *gorm.DB, txOpts *sql.TxOptions, retry RetryOptions, operation func(tx *gorm.DB) error) error {
	if retry.MaxRetries == 0 {
		retry.MaxRetries = DefaultMaxRetries
	}

	if retry.MaxRetries < 0 {
		retry.MaxRetries = math.MaxInt32
	}

	if retry.BackOff == nil {
		retry.BackOff = newSimpleExponentialBackOff().NextBackOff
	}

	if retry.IsRetryable == nil {
		retry.IsRetryable = DefaultIsRetryable
	}

	if retry.Sleep == nil {
		retry.Sleep = time.Sleep
	}

	var err error
	for i := 0; i < retry.MaxRetries; i++ {
		err = TransactWithOptions(db, txOpts, operation)
		if err == nil {
			return nil
		}
		if !retry.IsRetryable(err) {
			return err
		}
		retry.Sleep(retry.BackOff())
	}
	return err
}

// RollbackError is the error returned if the transaction operation returned an error, and the rollback automatically
// attempted also returns an error.
type RollbackError struct {
	// The original error that the operation returned.
	OriginalErr error
	// The error returned by sql.Tx.Rollback()
	Err error
}

// Unwrap returns the OriginalErr.
func (r *RollbackError) Unwrap() error {
	return r.OriginalErr
}

// Cause returns the OriginalErr.
func (r *RollbackError) Cause() error {
	return r.Unwrap()
}

// Error returns a formatted error message containing both the OriginalErr and RollbackError.
func (r *RollbackError) Error() string {
	return fmt.Sprintf("error rolling back failed transaction: %v, original transaction error: %v", r.Err, r.OriginalErr)
}

// RetryOptions controls how TransactWithRetry behaves.
type RetryOptions struct {
	// MaxRetries configures how many attempts will be made to complete the operation when a retryable error is
	// encountered. The default is DefaultMaxRetries. If set to a negative number, math.MaxInt32 attempts will be made.
	MaxRetries int
	// BackOff is called on each retry, and should return a time.Duration indicating how long to wait before the next
	// attempt. The default is an exponential backoff based on the values of DefaultInitialBackOff, DefaultMaxBackOff,
	// and DefaultBackOffFactor. If a negative Duration is returned by NextBackOff(), retries will be aborted.
	//
	// Most backoff implementations are compatible, including github.com/cenkalti/backoff and
	// github.com/jpillora/backoff.
	BackOff func() time.Duration
	// IsRetryable determines whether the error from the operation should be retried. Return true to retry.
	IsRetryable func(err error) bool
	// Sleep is an optional value to be used for mocking out time.Sleep() for testing. If set, backoff wait
	// will use this function instead of time.Sleep().
	Sleep func(duration time.Duration)
}

// nolint:deadcode
func alwaysRetryable(error) bool {
	return true
}

func retryIfApplicable(err error) bool {
	dbError := &dbErrors.DBError{}
	if errors.As(err, &dbError) {
		return dbError.Retryable
	}

	return false
}
