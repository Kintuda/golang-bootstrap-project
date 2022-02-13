package config

import (
	"github.com/Netflix/go-env"
	_ "github.com/joho/godotenv/autoload"
)

type ApplicationConfig struct {
	Database DatabaseConfig
	Runtime  RuntimeConfig
}

type RuntimeConfig struct {
	Env      string `env:"ENV,default=development"`
	HttpPort string `env:"HTTP_PORT,default=:3000"`
}

type DatabaseConfig struct {
	Dns   string `env:"DATABASE_URL,required=true"`
	Debug bool   `env:"DATABASE_DEBUG,default=false"`
}

func LoadConfigFromEnv() (*ApplicationConfig, error) {
	var cfg ApplicationConfig

	if _, err := env.UnmarshalFromEnviron(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
