package main

import (
	"github.com/gin-gonic/gin"
)

func (s *server) loadRoutes() *gin.Engine {
	r := gin.Default() // includes logger and recovery middleware

	api := r.Group("/api")

	//users
	users := api.Group("/users")
	users.POST("/", s.h.CreateUser)
	users.POST("/verify", s.h.VerifyUser)
	users.POST("/verify/new", s.h.RequestVerification)
	users.GET("/:id", s.h.GetUser)
	users.DELETE("/:id", s.h.DeleteUser)

	//users
	files := api.Group("/files")
	files.POST("/", s.h.UploadFile)
	files.GET("/users/:id", s.h.GetUserFiles)
	files.DELETE("/:id", s.h.DeleteFile)

	return r
}