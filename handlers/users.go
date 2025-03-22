package handlers

import "net/http"

func (h *AppHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

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
