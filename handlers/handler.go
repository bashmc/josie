package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/shcmd/josie/services"
)

// TODO: remove global variable
var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

type Handler struct {
	us *services.UserService
}

func NewHandler(us *services.UserService) *Handler {
	return &Handler{
		us: us,
	}
}
