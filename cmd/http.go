package main

import (
	"example/tm/authservice/config"
	"example/tm/authservice/internal/adapter/database/postgres"
	"example/tm/authservice/internal/adapter/router"
	"example/tm/authservice/pkg/logger"
	"log/slog"
	"os"
)

func main() {
	// init logger
	logger.Init()

	// load configs
	c, configErr := config.New()
	if configErr != nil {
		slog.Error("Error loading configuration:", configErr)
		os.Exit(1)
	}

	// db connection
	db, dbErr := postgres.New(*c.PostgresConfig)
	if dbErr != nil {
		slog.Error("Error on database connection", dbErr)
		os.Exit(1)
	}

	// init modules

	r := router.InitRouter()

	err := r.Serve(c.HttpConfig.Host + ":" + c.HttpConfig.Port)
	if err != nil {
		slog.Error("Error on starting error", err)
		os.Exit(1)
	}
}
