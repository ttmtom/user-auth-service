package router

import (
	"example/tm/authservice/config"
	"example/tm/authservice/internal/adapter/handler"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"log/slog"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	config *config.HttpConfig,
	userHandler *handler.UserHandler,
) *Router {
	router := gin.New()
	router.Use(sloggin.New(slog.Default()), gin.Recovery())

	v1 := router.Group("/v1")
	{
		user := v1.Group("/users")
		{
			user.POST("/", userHandler.Reg)
		}
	}

	return &Router{router}
}

func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
