package handlers

import (
	"errors"
	"net/http"

	"github.com/fatcmd/josie/models"
	"github.com/fatcmd/josie/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8,max=20"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse request body"})
		return
	}

	user, err := h.us.CreateUser(c.Request.Context(), input.Name, input.Email, input.Password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateUser) {
			c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (h *Handler) VerifyUser(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required,email"`
		Code  string `json:"code" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse request body"})
		return
	}

	user, err := h.us.VerifyUser(c.Request.Context(), input.Code, input.Email)
	if err != nil {
		if errors.Is(err, services.ErrInvalidToken) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *Handler) RequestVerification(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse request body"})
		return
	}

	err := h.us.ResendOTP(c.Request.Context(), input.Email)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse request body"})
		return
	}

	session, err := h.us.NewSession(c.Request.Context(), input.Email, input.Password)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	c.JSON(http.StatusOK, session)
}

func (h *Handler) GetUser(c *gin.Context) {
	userId := c.Param("id")
	if _, err := uuid.Parse(userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	user, err := h.us.FetchUser(c.Request.Context(), uuid.MustParse(userId))
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	if _, err := uuid.Parse(userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	err := h.us.DeleteUser(c.Request.Context(), userId)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user successfully deleted"})
}
