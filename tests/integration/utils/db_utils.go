package utils

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/NicklasWallgren/go-template/infrastructure/database"

	"github.com/NicklasWallgren/go-template/config"
)

func CreateDatabase(databaseName string, appConfig *config.AppConfig) (string, error) {
	db, err := sql.Open(appConfig.Database.Dialect, appConfig.Database.ToDatabaseDsn())
	if err != nil {
		return "", err
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + databaseName) // TODO, db dialect/driver specific
	return databaseName, err
}

func DropDatabase(databaseName string, appConfig *config.AppConfig) (string, error) {
	db, err := sql.Open(appConfig.Database.Dialect, appConfig.Database.ToDatabaseDsn())
	if err != nil {
		return "", err
	}

	_, err = db.Exec("DROP DATABASE IF EXISTS " + databaseName) // TODO, db dialect/driver specific
	return databaseName, err
}

func SeedDatabase() {}

func TruncateDatabase(db database.Database, config *config.AppConfig) error {
	// Truncates all tables except the goose specific table
	sql := fmt.Sprintf(`SET @tables = NULL;
			SELECT GROUP_CONCAT(table_schema, '.', table_name) INTO @tables FROM information_schema.tables
  			WHERE table_schema = '%s' AND table_name <> 'goose_db_version';
			SET @tables = CONCAT('TRUNCATE TABLE ', @tables);
			PREPARE stmt1 FROM @tables;
			EXECUTE stmt1;
			DEALLOCATE PREPARE stmt1;`, config.Database.Name)

	return db.Exec(sql).Error
}

func ToDatabaseName(testFuncName string) string {
	names := strings.Split(testFuncName, "/")

	return names[len(names)-1]
}
