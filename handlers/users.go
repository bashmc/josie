package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/shcmd/josie/models"
	"github.com/shcmd/josie/services"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	writeJson(w, http.StatusCreated, Map{"user": user})
}

func (h *Handler) VerifyUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email string `json:"email" validate:"required,email"`
		Code  string `json:"code" validate:"required"`
	}

	err := readJson(w, r, &input)
	if err != nil {
		writeJson(w, http.StatusBadRequest, Map{"message": "failed to parse request body"})
		return
	}

	err = validate.Struct(input)
	if err != nil {
		writeJson(w, http.StatusBadRequest, Map{"message": err.Error()})
		return
	}

	user, err := h.us.VerifyUser(r.Context(), input.Code, input.Email)
	if err != nil {
		if errors.Is(err, services.ErrInvalidToken) {
			writeJson(w, http.StatusBadRequest, Map{"message": err.Error()})
			return
		}

		serverError(w)
		return
	}

	writeJson(w, http.StatusOK, Map{"user": user})

}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	err := readJson(w, r, &input)
	if err != nil {
		writeJson(w, http.StatusBadRequest, Map{"message": "failed to parse request body"})
		return
	}
	err = validate.Struct(input)
	if err != nil {
		writeJson(w, http.StatusBadRequest, Map{"message": err.Error()})
		return
	}

	session, err := h.us.NewSession(r.Context(), input.Email, input.Password)
	if err != nil {
		//TODO: handle error
		return
	}

	writeJson(w, http.StatusOK, session)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")
	err := validate.Var(userId, "uuid")
	if err != nil {
		writeJson(w, http.StatusBadRequest, Map{"message": "invalid id"})
		return
	}

	user, err := h.us.FetchUser(r.Context(), uuid.MustParse(userId))
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			writeJson(w, http.StatusNotFound, Map{"message": err.Error()})
			return
		}

		serverError(w)
		return
	}

	writeJson(w, http.StatusOK, user)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")
	err := validate.Var(userId, "uuid")
	if err != nil {
		writeJson(w, http.StatusBadRequest, Map{"message": "invalid id"})
		return
	}

	err = h.us.DeleteUser(r.Context(), userId)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			writeJson(w, http.StatusNotFound, Map{"message": err.Error()})
			return
		}

		serverError(w)
		return
	}

	writeJson(w, http.StatusOK, Map{"message": "user successfully deleted"})
}
