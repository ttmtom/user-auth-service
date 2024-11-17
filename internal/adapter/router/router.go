package router

import (
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"log/slog"
)

type Router struct {
	*gin.Engine
}

func InitRouter() *Router {
	router := gin.New()
	router.Use(sloggin.New(slog.Default()), gin.Recovery())

	router.Group("/v1")
	{

	}

	return &Router{router}
}

func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
