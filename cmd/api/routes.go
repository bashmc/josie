package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *server) loadRoutes() *chi.Mux {
	r := chi.NewMux()

	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		// users
		r.Route("/users", func(r chi.Router) {
			r.Post("/", s.h.CreateUser)
			r.Get("/{id}", s.h.GetUser)
			r.Delete("/{id}", s.h.DeleteUser)
		})
		// files
		r.Route("/files", func(r chi.Router) {
			r.Post("/", s.h.UploadFile)
			r.Get("/users/{user_id}", s.h.GetUserFiles)
			r.Delete("/{file_id}", s.h.DeleteFile)
		})
	})

	return r
}
