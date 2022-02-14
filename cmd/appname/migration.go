package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Kintuda/golang-bootstrap-project/app"
	"github.com/Kintuda/golang-bootstrap-project/config"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

var (
	MigrationCmd = &cobra.Command{
		Use: "migration",
	}
	migrateUp = &cobra.Command{
		Use:  "up",
		RunE: migrationUp,
	}
	migrateDown = &cobra.Command{
		Use:  "down",
		RunE: migrationDown,
	}
	createMigrationFile = &cobra.Command{
		RunE: createMigration,
		Use:  "create",
		Args: cobra.MinimumNArgs(1),
	}
	migrationDir string
)

func init() {
	MigrationCmd.PersistentFlags().StringVarP(&migrationDir, "migration-dir", "m", "", "directory that holds the migration files")
	MigrationCmd.AddCommand(migrateUp, migrateDown, createMigrationFile)
}

func createMigration(cmd *cobra.Command, args []string) error {
	var migrationName = strings.Join(args, "_")
	var revision int64 = 1
	var initialMigration = false

	migrationDirectory := filepath.Clean(migrationDir)
	migrationFiles, err := filepath.Glob(filepath.Join(migrationDirectory, "*.sql"))

	if err != nil {
		return err
	}

	if len(migrationFiles) <= 0 {
		initialMigration = true
	}

	if !initialMigration {
		lastMigration := migrationFiles[len(migrationFiles)-1]
		baseFileName := filepath.Base(lastMigration)
		fileRevision := strings.Index(baseFileName, "_")

		revisionString := baseFileName[0:fileRevision]
		nextRevision, err := strconv.ParseInt(revisionString, 10, 64)

		if err != nil {
			return err
		}

		nextRevision++
		revision = int64(nextRevision)
		fmt.Println(revision)
	}

	migrationName = strings.ReplaceAll(strings.ToLower(migrationName), " ", "_")

	for _, direction := range []string{"up", "down"} {
		fmt.Println(strconv.FormatInt(revision, 10))
		basename := fmt.Sprintf("%s_%s.%s%s", strconv.FormatInt(revision, 10), migrationName, direction, ".sql")
		filename := filepath.Join(migrationDirectory, basename)

		if err = createFile(filename); err != nil {
			return err
		}

		absPath, _ := filepath.Abs(filename)
		fmt.Println(absPath)
	}

	return nil
}

func createFile(filename string) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)

	if err != nil {
		return err
	}

	return f.Close()
}

func migrationUp(cmd *cobra.Command, arg []string) error {
	cfg, err := config.LoadConfigFromEnv()

	if err != nil {
		return err
	}

	m, err := app.NewMigrator(&cfg.Database, migrationDir)

	if err != nil {
		return err
	}

	defer func() {
		if _, err := m.Engine.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := m.Engine.Up(); err != nil {
		return err
	}

	fmt.Println("Migration up ran successfully")

	return nil
}

func migrationDown(cmd *cobra.Command, arg []string) error {
	cfg, err := config.LoadConfigFromEnv()

	if err != nil {
		return err
	}

	m, err := app.NewMigrator(&cfg.Database, migrationDir)

	if err != nil {
		return err
	}

	defer func() {
		if _, err := m.Engine.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := m.Engine.Down(); err != nil {
		return err
	}

	fmt.Println("Migration down ran successfully")

	return nil
}
