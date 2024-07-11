package user

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Route(h *server.Hertz) {
	userHandler := NewHandler()
	h.POST("/signup", userHandler.Signup)
	h.POST("/login", userHandler.Login)
}
