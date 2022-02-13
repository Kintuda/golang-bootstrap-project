package cmd

import (
	"github.com/Kintuda/golang-bootstrap-project/app"
	config "github.com/Kintuda/golang-bootstrap-project/config"
	"github.com/Kintuda/golang-bootstrap-project/db"
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "start appname API",
	RunE:  startCmd,
}

func startCmd(cmd *cobra.Command, arg []string) error {
	cfg, err := config.LoadConfigFromEnv()

	if err != nil {
		return err
	}

	db, err := db.NewDatabaseConnection(&cfg.Database)

	if err != nil {
		return err
	}

	server := app.NewServer(db, cfg)
	app.NewAPI(server)

	if err != nil {
		return err
	}

	err = server.Init()

	return err
}
