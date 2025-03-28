package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/xpmc/split/models"
)

func (h *AppHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := readJson(w, r, &input)
	if err != nil {
		slog.Error("failed to read request body", "error", err)
		writeJson(w, http.StatusBadRequest, Map{"message": "failed to parse request body"})
		return
	}

	user, err := h.us.CreateUser(r.Context(), input.Name, input.Email, input.Password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateUser) {
			writeJson(w, http.StatusConflict, Map{"message": err.Error()})
			return
		}

		sendServerError(w)
		return
	}

	writeJson(w, http.StatusCreated, user)
}

func (h *AppHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")

	user, err := h.us.FetchUser(r.Context(), userId)
	if err != nil {
		writeJson(w, http.StatusNotFound, Map{"error": "User not found"})
		return
	}

	writeJson(w, http.StatusOK, user)
}
