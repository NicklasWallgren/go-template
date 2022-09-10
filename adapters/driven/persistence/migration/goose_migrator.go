package migration

import (
	"database/sql"
	"fmt"
	"io/fs"

	"github.com/NicklasWallgren/go-template/adapters/driven/persistence"

	"github.com/NicklasWallgren/go-template/config"
	"github.com/pressly/goose/v3"
)

type GooseMigrator struct {
	db                     persistence.Database
	filesystem             fs.FS
	migrationDirectoryPath string
}

func NewGooseMigrator(db persistence.Database, config *config.AppConfig) (Migrator, error) {
	goose.SetBaseFS(config.Assets.EmbedMigrations)

	if err := goose.SetDialect(config.Database.Driver); err != nil {
		return nil, err
	}

	return &GooseMigrator{
		db:                     db,
		filesystem:             config.Assets.EmbedMigrations,
		migrationDirectoryPath: config.Database.MigrationDirectory,
	}, nil
}

func (m GooseMigrator) Up() error {
	db, err := m.dbConnection()
	if err != nil {
		return err
	}

	if err := goose.Up(db, m.migrationDirectoryPath, m.optionFuncs()...); err != nil {
		return fmt.Errorf("could not apply the available migrations. %w", err)
	}

	return nil
}

func (m GooseMigrator) Down() error {
	// TODO implement me
	panic("implement me")

	// down
	// downTo
} // nolint:wsl

func (m GooseMigrator) Create(name string) error {
	db, err := m.dbConnection()
	if err != nil {
		return err
	}

	if err := goose.Create(db, m.migrationDirectoryPath, name, "sql"); err != nil {
		return err
	}

	return nil
}

func (m GooseMigrator) Fix() error {
	// TODO implement me
	panic("implement me")
}

func (m GooseMigrator) Status() error {
	db, err := m.dbConnection()
	if err != nil {
		return err
	}

	if err := goose.Status(db, m.migrationDirectoryPath, m.optionFuncs()...); err != nil {
		return err
	}

	return nil
}

func (m GooseMigrator) dbConnection() (*sql.DB, error) {
	db, err := m.db.DB.DB()
	if err != nil {
		return nil, fmt.Errorf("could not establish a connection to the database. %w", err)
	}

	return db, nil
}

func (m GooseMigrator) optionFuncs() []goose.OptionsFunc {
	return []goose.OptionsFunc{
		goose.WithAllowMissing(), // TODO, option to control?
		// goose.WithNoVersioning(), // TODO, should be active by default?
	}
}
