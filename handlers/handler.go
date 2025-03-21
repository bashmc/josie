package handlers

import "github.com/u88x/split/services"

type AppHandler struct {
	us *services.UserService
}

func NewAppHandler(us *services.UserService) *AppHandler {
	return &AppHandler{
		us: us,
	}
}
