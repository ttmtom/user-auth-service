package main

import (
	"example/tm/authservice/config"
	"example/tm/authservice/internal/adapter/database/postgres"
	"example/tm/authservice/internal/adapter/database/postgres/repository"
	"example/tm/authservice/internal/adapter/handler"
	"example/tm/authservice/internal/adapter/router"
	"example/tm/authservice/internal/core/service"
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
	userRepository := repository.NewUserRepository(db.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	r := router.NewRouter(
		c.HttpConfig,
		userHandler,
	)

	err := r.Serve(c.HttpConfig.Host + ":" + c.HttpConfig.Port)
	if err != nil {
		slog.Error("Error on starting error", err)
		os.Exit(1)
	}
}
