package config

type ApplicationConfig struct {
	Database DatabaseConfig
	Runtime  RuntimeConfig
}

type RuntimeConfig struct {
	Env string `env:"ENV" validate:"required"`
}

type DatabaseConfig struct {
	Dns string `env:"DATABASE_URL" validate:"required"`
}
