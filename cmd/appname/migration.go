package cmd

import (
	"database/sql"
	"errors"

	"github.com/Kintuda/golang-bootstrap-project/config"
	env "github.com/Netflix/go-env"
	"github.com/go-playground/validator/v10"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var MigrationCommand = &cobra.Command{
	Use: "migration",
}

var CreateMigrationCommand = &cobra.Command{
	Use:  "create",
	RunE: CreateMigration,
}

func Init() {
	migrationCommand.AddCommand(CreateMigration())
}

func CreateMigration(cmd *cobra.Command, arg []string) error {
	var cfg config.DatabaseConfig
	var err error

	if err := godotenv.Load(); err != nil {
		return errors.New("error while loading .env file")
	}

	_, err = env.UnmarshalFromEnviron(&cfg)

	if err != nil {
		return err
	}

	validate := validator.New()

	if err := validate.Struct(cfg); err != nil {
		return err
	}

	db, err := sql.Open("postgres", cfg.Dns)

	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		return nil
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///db/migrations",
		"postgres", driver)

	if err != nil {
		return nil
	}

	m.Steps(2)

	return nil
}
