package api

import (
	"encoding/json"
	"net/http"
)

func write(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if body != nil {
		json.NewEncoder(w).Encode(body)
	}
}

func OKResponse(w http.ResponseWriter, data any) {
	write(w, http.StatusOK, data)
}

func ErrorResponse(w http.ResponseWriter, status int, message string) {
	write(w, status, map[string]string{"error": message})
}
