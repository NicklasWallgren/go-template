package drivers

import (
	"errors"

	dbErrors "github.com/NicklasWallgren/go-template/adapters/driver/persistence/errors"
	"github.com/go-sql-driver/mysql"
)

// To ensure that MySQLDriver implements the Driver interface.
var _ Driver = (*MySQLDriver)(nil)

const (
	mysqlDuplicateEntry   = 1062
	mysqlLockWaitTimeout  = 1205
	mysqlIncorrectInteger = 1366 // nolint:deadcode, varcheck
)

type MySQLDriver struct{}

func (m MySQLDriver) ConvertError(driverError error) error {
	mysqlError := &mysql.MySQLError{}
	if !errors.As(driverError, &mysqlError) {
		return dbErrors.NewDBError(dbErrors.WithRetryableAndError(false, driverError))
	}

	switch mysqlError.Number {
	case mysqlDuplicateEntry:
		// TODO, might be useful with retry?
		return dbErrors.NewDBError(
			dbErrors.WithRetryableAndTypeAndError(false, dbErrors.DuplicateEntry, mysqlError))
	case mysqlLockWaitTimeout:
		return dbErrors.NewDBError(
			dbErrors.WithRetryableAndTypeAndError(true, dbErrors.LockWaitTimeout, mysqlError))
	}

	return dbErrors.NewDBError(dbErrors.WithError(mysqlError))
}
