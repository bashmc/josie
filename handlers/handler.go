package handlers

import (
	"github.com/fatcmd/josie/services"
)

type Handler struct {
	us *services.UserService
}

func NewHandler(us *services.UserService) *Handler {
	return &Handler{
		us: us,
	}
}
