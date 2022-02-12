package config

import (
	"errors"

	"github.com/Netflix/go-env"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

var (
	errInvalidEnvFile = errors.New("error while loading .env file")
)

type ApplicationConfig struct {
	Database DatabaseConfig
	Runtime  RuntimeConfig
}

type RuntimeConfig struct {
	Env      string `env:"ENV" validate:"required"`
	HttpPort string `env:"HTTP_PORT" validate:"required"`
}

type DatabaseConfig struct {
	Dns   string `env:"DATABASE_URL" validate:"required"`
	Debug bool   `env:"DATABASE_DEBUG"`
}

func LoadDatabaseCredentialsFromEnv() (*DatabaseConfig, error) {
	var cfg DatabaseConfig

	if err := godotenv.Load(); err != nil {
		return nil, errInvalidEnvFile
	}

	if _, err := env.UnmarshalFromEnviron(&cfg); err != nil {
		return nil, err
	}

	validate := validator.New()

	if err := validate.Struct(cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
