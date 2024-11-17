package config

import (
	"github.com/joho/godotenv"
	"log/slog"
)

type Config struct {
	AppConfig      *AppConfig
	HttpConfig     *HttpConfig
	PostgresConfig *PostgresConfig
}

func New() (*Config, error) {
	envFile, envFileLoadingErr := godotenv.Read(".env")
	if envFileLoadingErr != nil {
		return nil, envFileLoadingErr
	}

	slog.Info("Env File load successfully", "file", envFile)

	appConfig := LoadAppConfig(envFile)
	httpConfig := LoadHttpConfig(envFile)
	postgresConfig := LoadPostgresConfig(envFile)

	return &Config{
		appConfig,
		httpConfig,
		postgresConfig,
	}, nil
}

func NewMigration() (*PostgresConfig, error) {
	envFile, envFileLoadingErr := godotenv.Read(".env")
	if envFileLoadingErr != nil {
		return nil, envFileLoadingErr
	}

	postgresConfig := LoadPostgresConfig(envFile)

	return postgresConfig, nil
}
