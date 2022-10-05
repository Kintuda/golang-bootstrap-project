package cmd

import (
	"github.com/Kintuda/golang-bootstrap-project/pkg/config"
	"github.com/Kintuda/golang-bootstrap-project/pkg/http"
	"github.com/Kintuda/golang-bootstrap-project/pkg/postgres"
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "start appname API",
	RunE:  startCmd,
}

func startCmd(cmd *cobra.Command, arg []string) error {
	cfg := &config.AppConfig{}

	if err := config.LoadConfigFromEnv(cfg); err != nil {
		return err
	}

	db, err := postgres.NewDatabaseConnection(cfg.PostgresDns)

	if err != nil {
		return err
	}

	router, err := http.NewRouter(cfg, db)

	if err != nil {
		return err
	}

	server := http.NewServer(router, cfg.HttpPort)

	err = server.Init()
	return err
}
