package handlers

import "github.com/tcpbot/split/services"

type AppHandler struct {
	us *services.UserService
}

func NewAppHandler(us *services.UserService) *AppHandler {
	return &AppHandler{
		us: us,
	}
}
