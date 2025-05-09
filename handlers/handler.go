package handlers

import (
	"github.com/topkobie/josie/services"
)

type Handler struct {
	us *services.UserService
}

func NewHandler(us *services.UserService) *Handler {
	return &Handler{
		us: us,
	}
}
