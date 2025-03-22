package main

import (
	"net/http"

	"github.com/xpmc/split/handlers"
)

type Server struct {
	h *handlers.AppHandler
}

func NewServer(handler *handlers.AppHandler) *Server {
	return &Server{
		h: handler,
	}
}

func (s *Server) start() error {
	srv := http.Server{
		Addr: ":7000",
		Handler: s.loadRoutes(),
	}

	return srv.ListenAndServe()
}
