package drivers

import (
	"errors"

	dbErrors "github.com/NicklasWallgren/go-template/adapters/driver/persistence/errors"
	"github.com/go-sql-driver/mysql"
)

// To ensure that MySQLDriver implements the Driver interface
var _ Driver = (*MySQLDriver)(nil)

const (
	DuplicateEntry   = 1062
	LockWaitTimeout  = 1205
	IncorrectInteger = 1366
)

type MySQLDriver struct{}

func (m MySQLDriver) ConvertError(driverError error) error {
	mysqlError := &mysql.MySQLError{}
	if !errors.As(driverError, &mysqlError) {
		return dbErrors.NewDBError(dbErrors.WithRetryableAndError(false, driverError))
	}

	switch mysqlError.Number {
	case DuplicateEntry:
		return dbErrors.NewDBError(dbErrors.WithRetryableAndTypeAndError(false, DuplicateEntry, mysqlError)) // TODO, might be useful with retry?
	case LockWaitTimeout:
		return dbErrors.NewDBError(dbErrors.WithRetryableAndTypeAndError(true, LockWaitTimeout, mysqlError))
	}

	return dbErrors.NewDBError(dbErrors.WithRetryableAndError(false, mysqlError))
}
