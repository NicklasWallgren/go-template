package cmd

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/migration"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

type MigrationCommand struct{}

func NewMigrationCommand() Command {
	return &MigrationCommand{}
}

func (m MigrationCommand) Use() string {
	return "migrate"
}

func (m MigrationCommand) Short() string {
	return "Handle database migration"
}

func (m MigrationCommand) Setup(cmd *cobra.Command) {
	cmd.Flags().Bool("up", false, "Applies the migration files")
	cmd.Flags().String("create", "", "Creates a new migration file")
}

func (m MigrationCommand) Run(cmd *cobra.Command) CommandRunner {
	return func(migrator migration.Migrator) {
		upFlag, _ := cmd.Flags().GetBool("up")
		createFlag, _ := cmd.Flags().GetString("create")

		if upFlag {
			migrate(migrator)
		} else if createFlag != "" {
			create(migrator, createFlag)
		}

		// TODO, handle unknown flag
	} // nolint: wsl
}

func migrate(migrator migration.Migrator) {
	if err := migrator.Up(); err != nil {
		// TODO, handle error properly
		panic(err)
	}
}

func create(migrator migration.Migrator, filename string) {
	if err := migrator.Create(filename); err != nil {
		// TODO, handle error properly
		panic(err)
	}
}
