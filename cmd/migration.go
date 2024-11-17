package main

import (
	"example/tm/authservice/config"
	"example/tm/authservice/pkg/logger"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"log/slog"
	"os"
)

func main() {
	// init logger
	logger.Init()

	argType := "up"
	if len(os.Args) > 1 {
		argType = os.Args[1]
	}
	slog.Info(argType)

	pgConfig, err := config.NewMigration()
	if err != nil {
		slog.Error("Error on loading DB config", err)
		os.Exit(1)
	}

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		pgConfig.DatabaseUser,
		pgConfig.DatabasePassword,
		pgConfig.PostgresHost,
		pgConfig.PostgresPort,
		pgConfig.PostgresDb,
	)
	sourceURL := "file://migration"

	m, err := migrate.New(
		sourceURL,
		databaseURL,
	)
	if err != nil {
		slog.Error("Error on DB connection", err)
		os.Exit(1)
	}

	if argType == "up" {
		err = m.Up()
	} else if argType == "down" {
		err = m.Steps(-1)
	}

	if err != nil {
		slog.Error(fmt.Sprintf("Error on migration %s", argType), err)
		os.Exit(1)
	}
	slog.Info(fmt.Sprintf("Migration %s completed successfully", argType))

}
