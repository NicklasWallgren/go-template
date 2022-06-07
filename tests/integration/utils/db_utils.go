package utils

import (
	"database/sql"
	"strings"

	_ "github.com/lib/pq"

	"github.com/NicklasWallgren/go-template/infrastructure/database"

	"github.com/NicklasWallgren/go-template/config"
)

func CreateDatabase(databaseName string, appConfig *config.AppConfig) (string, error) {
	db, err := sql.Open(appConfig.Database.Driver, database.DSN(appConfig.Database))
	if err != nil {
		return "", err
	}

	_, err = db.Exec("DROP DATABASE IF EXISTS " + databaseName) // TODO, db dialect/driver specific
	_, err = db.Exec("CREATE DATABASE " + databaseName)         // TODO, db dialect/driver specific
	return databaseName, err
}

func DropDatabase(databaseName string, appConfig *config.AppConfig) (string, error) {
	db, err := sql.Open(appConfig.Database.Driver, database.DSN(appConfig.Database))
	if err != nil {
		return "", err
	}

	_, err = db.Exec("DROP DATABASE IF EXISTS " + databaseName) // TODO, db dialect/driver specific
	return databaseName, err
}

func SeedDatabase() {}

func TruncateDatabase(db database.Database, config *config.AppConfig) error {
	// Truncates all tables except the goose specific table
	sql := "call truncate_tables('go_template_test')"

	return db.Exec(sql).Error
}

func ToDatabaseName(testFuncName string) string {
	names := strings.Split(testFuncName, "/")

	return strings.ToLower(names[len(names)-1])
}
