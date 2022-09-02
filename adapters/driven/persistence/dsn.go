package persistence

import (
	"fmt"

	"github.com/NicklasWallgren/go-template/config"
)

func DSN(config *config.Database) string {
	switch config.Driver {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&multiStatements=true",
			config.User, config.Password, config.Host, config.Port, config.Name)
	case "postgres":
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&lock_timeout=5000ms&TimeZone=UTC",
			config.User, config.Password, config.Host, config.Port, config.Name)
	}

	return ""
}
