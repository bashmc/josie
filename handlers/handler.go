package handlers

import (
	"github.com/gitsnack/josie/services"
)

type Handler struct {
	us *services.UserService
}

func NewHandler(us *services.UserService) *Handler {
	return &Handler{
		us: us,
	}
}
