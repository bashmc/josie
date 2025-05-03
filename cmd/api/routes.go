package main

import "net/http"

func (s *server) loadRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// users
	mux.HandleFunc("POST /users", s.h.CreateUser)
	mux.HandleFunc("GET /users/{id}", s.h.GetUser)
	mux.HandleFunc("DELETE /users/{id}", s.h.DeleteUser)

	// files
	mux.HandleFunc("POST /files", s.h.UploadFile)
	mux.HandleFunc("GET /files", s.h.GetUserFiles)
	mux.HandleFunc("DELETE /files", s.h.DeleteFile)

	return mux
}
