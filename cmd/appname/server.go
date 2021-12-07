package cmd

import (
	"errors"

	"github.com/Kintuda/golang-bootstrap-project/app"
	"github.com/Kintuda/golang-bootstrap-project/config"
	"github.com/Kintuda/golang-bootstrap-project/db"
	v1 "github.com/Kintuda/golang-bootstrap-project/v1"
	env "github.com/Netflix/go-env"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "start appname API",
	RunE:  startCmd,
}

func startCmd(cmd *cobra.Command, arg []string) error {
	var cfg config.ApplicationConfig
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

	db, err := db.NewDatabaseConnection(&cfg.Database)

	if err != nil {
		return err
	}

	server := app.NewServer(db, &cfg)
	v1.NewAPI(server)

	if err != nil {
		return err
	}

	err = server.Init()

	return err
}
