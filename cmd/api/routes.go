package main

import "net/http"

func (s *Server) loadRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", s.h.CreateUser)
	mux.HandleFunc("GET /users/{id}", s.h.GetUser)
	mux.HandleFunc("DELETE /users/{id}", s.h.DeleteUser)

	return mux
}
