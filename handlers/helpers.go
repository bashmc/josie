package handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type Map map[string]any

func readJson(w http.ResponseWriter, r *http.Request, dest any) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return errors.New("Content-Type of 'application/json' expected")
	}

	maxBytes := 1048576 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(dest)
	if err != nil {
		return err
	}

	return nil
}

func writeJson(w http.ResponseWriter, statusCode int, data any, headers ...http.Header) {
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			slog.Error("failed to encode response", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Map{"error": "Internal server error"})
		}
	}
}

func serverError(w http.ResponseWriter, headers ...http.Header) {
	writeJson(w, http.StatusInternalServerError, Map{"message": "the server could not process your request"}, headers...)
}
