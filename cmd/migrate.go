package cmd

import (
	"github.com/NicklasWallgren/go-template/infrastructure/cli"
	"github.com/NicklasWallgren/go-template/infrastructure/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

type MigrationCommand struct{}

func NewMigrationCommand() cli.Command {
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

func (m MigrationCommand) Run(cmd *cobra.Command) cli.CommandRunner {
	return func(migrator database.Migrator) {
		upFlag, _ := cmd.Flags().GetBool("up")
		createFlag, _ := cmd.Flags().GetString("create")

		if upFlag {
			migrate(migrator)
		} else if createFlag != "" {
			create(migrator, createFlag)
		}

		// TODO, handle unknown flag
	}
}

func migrate(migrator database.Migrator) {
	if err := migrator.Up(); err != nil {
		// TODO, handle error properly
		panic(err)
	}
}

func create(migrator database.Migrator, filename string) {
	if err := migrator.Create(filename); err != nil {
		// TODO, handle error properly
		panic(err)
	}
}
