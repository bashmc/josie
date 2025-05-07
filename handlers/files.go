package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UploadFile(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse form"})
		return
	}

	files := c.Request.MultipartForm.File

	for _, value := range files {
		for i := range value {
			fmt.Println(value[i].Filename)
		}
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteFile(c *gin.Context) {
	// Implement your delete logic here
	c.Status(http.StatusOK)
}

func (h *Handler) GetUserFiles(c *gin.Context) {
	// Implement your get files logic here
	c.Status(http.StatusOK)
}
