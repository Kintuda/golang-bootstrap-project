package cmd

import (
	"github.com/Kintuda/golang-bootstrap-project/pkg/config"
	"github.com/Kintuda/golang-bootstrap-project/pkg/http"
	"github.com/Kintuda/golang-bootstrap-project/pkg/postgres"
	"github.com/spf13/cobra"
)

func NewServerCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "server",
		Short: "start appname API",
	}

	listen := &cobra.Command{
		Use:  "listen",
		RunE: startCmd,
	}

	rootCmd.AddCommand(listen)

	return rootCmd
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
