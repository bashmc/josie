package main

import (
	"net/http"

	"github.com/shcmd/josie/handlers"
)

type server struct {
	h *handlers.Handler
}

func newserver(handler *handlers.Handler) *server {
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
