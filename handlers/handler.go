package handlers

import (
	"github.com/gitkobie/split/services"
	"github.com/go-playground/validator/v10"
)

// TODO: remove global variable
var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

type AppHandler struct {
	us *services.UserService
}

func NewAppHandler(us *services.UserService) *AppHandler {
	return &AppHandler{
		us: us,
	}
}
