package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gitkobie/split/models"
)

func (h *AppHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,gte=8,lte=20"`
	}

	err := readJson(w, r, &input)
	if err != nil {
		slog.Error("failed to read request body", "error", err)
		writeJson(w, http.StatusBadRequest, Map{"message": "failed to parse request body"})
		return
	}

	err = validate.Struct(input)
	if err != nil {
		writeJson(w, http.StatusBadRequest, Map{"message": err.Error()})
		return
	}

	user, err := h.us.CreateUser(r.Context(), input.Name, input.Email, input.Password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateUser) {
			writeJson(w, http.StatusConflict, Map{"message": err.Error()})
			return
		}

		serverError(w)
		return
	}

	writeJson(w, http.StatusCreated, user)
}

func (h *AppHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")
	err := validate.Var(userId, "uuid")
	if err != nil {
		writeJson(w, http.StatusBadRequest, Map{"message": "invalid id"})
		return
	}

	user, err := h.us.FetchUser(r.Context(), userId)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			writeJson(w, http.StatusNotFound, Map{"message": err.Error()})
			return
		}

		serverError(w)
		return
	}

	writeJson(w, http.StatusOK, user)
}

func (h *AppHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")
	err := validate.Var(userId, "uuid")
	if err != nil {
		writeJson(w, http.StatusBadRequest, Map{"message": "invalid id"})
		return
	}

	err = h.us.DeleteUser(r.Context(), userId)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			writeJson(w, http.StatusNotFound, Map{"message": err.Error()})
			return
		}

		serverError(w)
		return
	}

	writeJson(w, http.StatusOK, Map{"message": "user successfully deleted"})
}
