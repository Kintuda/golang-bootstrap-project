package app

import (
	"database/sql"
	"fmt"

	"github.com/Kintuda/golang-bootstrap-project/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrator struct {
	Engine        *migrate.Migrate
	Cfg           *config.DatabaseConfig
	MigrationPath string
}

func NewMigrator(cfg *config.DatabaseConfig, migrationPath string) (*Migrator, error) {
	db, err := sql.Open("postgres", cfg.Dns)

	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		return nil, err
	}

	migrationPath = fmt.Sprintf("file://%v", migrationPath)

	m, err := migrate.NewWithDatabaseInstance(
		migrationPath,
		"postgres", driver)

	if err != nil {
		return nil, err
	}

	return &Migrator{
		Engine:        m,
		Cfg:           cfg,
		MigrationPath: migrationPath,
	}, nil
}
