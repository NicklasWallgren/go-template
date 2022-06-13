package database

import (
	"errors"
	"fmt"
	"io/fs"

	"github.com/NicklasWallgren/go-template/config"
	"github.com/pressly/goose/v3"
)

type GooseMigrator struct {
	db                     Database
	filesystem             fs.FS
	migrationDirectoryPath string
}

// ErrMigrationUp is returned when the available migrations could not be applied
var ErrMigrationUp = errors.New("could not apply the available migrations")

func NewGooseMigrator(db Database, config *config.AppConfig) (Migrator, error) {
	goose.SetBaseFS(config.Assets.EmbedMigrations)

	if err := goose.SetDialect(config.Database.Dialect); err != nil {
		return nil, err
	}

	return &GooseMigrator{db: db, filesystem: config.Assets.EmbedMigrations, migrationDirectoryPath: config.Database.MigrationDirectory}, nil
}

func (m GooseMigrator) Up() error {
	db, err := m.db.DB.DB()
	if err != nil {
		return fmt.Errorf("could not establish a connection to the database. %w", err)
	}

	// TODO, option to control WithAllowMissing?

	if err := goose.Up(db, m.migrationDirectoryPath, goose.WithAllowMissing()); err != nil {
		return fmt.Errorf("could not apply the available migrations. %w", err)
	}

	return nil
}

func (m GooseMigrator) Down() error {
	// TODO implement me
	panic("implement me")

	// down
	// downTo

}

func (m GooseMigrator) Create(name string) error {
	db, err := m.db.DB.DB()
	if err != nil {
		return fmt.Errorf("could not establish a connection to the database. %w", err)
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
