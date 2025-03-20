package handlers

import (
	"encoding/json"
	"errors"
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

func writeJson(w http.ResponseWriter, statusCode int, data any, headers ...http.Header) error {
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	return encoder.Encode(data)
}
