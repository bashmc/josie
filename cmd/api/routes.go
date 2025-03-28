package main

import "net/http"

func (s *Server) loadRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", s.h.CreateUser)
	mux.HandleFunc("GET /users/:id", s.h.GetUser)

	return mux
}
