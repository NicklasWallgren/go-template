package drivers

import (
	"errors"

	"github.com/jackc/pgconn"

	dbErrors "github.com/NicklasWallgren/go-template/adapters/driver/persistence/errors"
)

// To ensure that PostgresDriver implements the Driver interface
var _ Driver = (*PostgresDriver)(nil)

const (
	postgresDuplicateEntry  = "23505"
	postgresLockWaitTimeout = "1205"
)

type PostgresDriver struct{}

func (m PostgresDriver) ConvertError(driverError error) error {
	postgresError := &pgconn.PgError{}
	if !errors.As(driverError, &postgresError) {
		return dbErrors.NewDBError(dbErrors.WithRetryableAndError(false, driverError))
	}

	switch postgresError.Code {
	case postgresDuplicateEntry:
		return dbErrors.NewDBError(dbErrors.WithRetryableAndTypeAndError(false, dbErrors.DuplicateEntry, postgresError)) // TODO, might be useful with retry?
	case postgresLockWaitTimeout:
		return dbErrors.NewDBError(dbErrors.WithRetryableAndTypeAndError(true, dbErrors.LockWaitTimeout, postgresError))
	}

	return dbErrors.NewDBError(dbErrors.WithRetryableAndError(false, postgresError))
}
