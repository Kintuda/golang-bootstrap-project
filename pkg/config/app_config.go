package config

type AppConfig struct {
	Env         string `mapstructure:"ENV"`
	HttpPort    string `mapstructure:"HTTP_PORT"`
	PostgresDns string `mapstructure:"POSTGRES_DNS"`
}
