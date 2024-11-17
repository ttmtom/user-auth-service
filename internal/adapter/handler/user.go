package handler

import (
	"example/tm/authservice/internal/core/service"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc,
	}
}

func (uh *UserHandler) Reg(ctx *gin.Context) {
	slog.Info("user reg", ctx)
}
