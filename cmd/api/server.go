package main

import (
	"net/http"

	"github.com/shcmd/split/handlers"
)

type server struct {
	h *handlers.AppHandler
}

func newserver(handler *handlers.AppHandler) *server {
	return &server{
		h: handler,
	}
}

func (s *server) start() error {
	srv := http.Server{
		Addr:    ":7000",
		Handler: s.loadRoutes(),
	}

	return srv.ListenAndServe()
}
