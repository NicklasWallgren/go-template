package errors

import "fmt"

type DBErrorType int

const (
	Generic DBErrorType = iota
	DuplicateEntry
	LockWaitTimeout
	IncorrectInteger
)

type DBError struct {
	Retryable bool
	err       error
	ErrorType DBErrorType
}

type DBErrorOption func(dbError *DBError)

func WithRetryableAndTypeAndError(retryable bool, dbErrorType DBErrorType, err error) DBErrorOption {
	return func(dbError *DBError) {
		dbError.Retryable = retryable
		dbError.ErrorType = dbErrorType
		dbError.err = err
	}
}

func WithRetryableAndError(retryable bool, err error) DBErrorOption {
	return func(dbError *DBError) {
		dbError.Retryable = retryable
		dbError.err = err
	}
}

func WithError(err error) DBErrorOption {
	return func(dbError *DBError) {
		dbError.err = err
	}
}

func NewDBError(options ...DBErrorOption) *DBError {
	err := &DBError{ErrorType: Generic}

	for _, option := range options {
		option(err)
	}

	return err
}

func (d DBError) Error() string {
	return fmt.Sprintf("persistence error. cause: %s", d.err.Error())
}
