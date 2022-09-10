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
	return "handleMigrate"
}

func (m MigrationCommand) Short() string {
	return "Handle database migration"
}

func (m MigrationCommand) Setup(cmd *cobra.Command) {
	cmd.Flags().Bool("up", false, "Applies the migration files")
	cmd.Flags().String("handleCreate", "", "Creates a new migration file")
	cmd.Flags().Bool("handleStatus", false, "Shows the migration handleStatus")
}

func (m MigrationCommand) Run(cmd *cobra.Command) CommandRunner {
	return func(migrator migration.Migrator) {
		upFlag, _ := cmd.Flags().GetBool("up")
		createFlag, _ := cmd.Flags().GetString("handleCreate")
		statusFlag, _ := cmd.Flags().GetBool("handleStatus")

		switch {
		case upFlag:
			handleMigrate(migrator)
		case createFlag != "":
			handleCreate(migrator, createFlag)
		case statusFlag:
			handleStatus(migrator)
		default:
			// TODO, handle unknown flag
		} // nolint: wsl
	} // nolint: wsl
}

func handleMigrate(migrator migration.Migrator) {
	if err := migrator.Up(); err != nil {
		// TODO, handle error properly
		panic(err)
	}
}

func handleCreate(migrator migration.Migrator, filename string) {
	if err := migrator.Create(filename); err != nil {
		// TODO, handle error properly
		panic(err)
	}
}

func handleStatus(migrator migration.Migrator) {
	if err := migrator.Status(); err != nil {
		// TODO, handle error properly
		panic(err)
	}
}
