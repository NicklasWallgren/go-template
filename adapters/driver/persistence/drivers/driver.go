package drivers

type Driver interface {
	ConvertError(driverError error) error
}

// TODO, support other drivers.
var drivers = map[string]Driver{
	"mysql":    &MySQLDriver{},
	"postgres": &PostgresDriver{},
}

func GetDriverOrNil(dialect string) Driver {
	driver, ok := drivers[dialect]
	if !ok {
		return nil
	}

	return driver
}
