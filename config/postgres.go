package config

type PostgresConfig struct {
	DatabaseUser     string
	DatabasePassword string
	PostgresHost     string
	PostgresPort     string
	PostgresDb       string
}

func LoadPostgresConfig(envFile map[string]string) *PostgresConfig {
	return &PostgresConfig{
		DatabaseUser:     envFile["POSTGRES_USER"],
		DatabasePassword: envFile["POSTGRES_PASSWORD"],
		PostgresHost:     envFile["POSTGRES_HOST"],
		PostgresPort:     envFile["POSTGRES_PORT"],
		PostgresDb:       envFile["POSTGRES_DB"],
	}
}
