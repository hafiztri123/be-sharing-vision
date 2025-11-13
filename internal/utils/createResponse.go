package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    any    `json:"data,omitempty"`
}

func NewJSONResponse(w http.ResponseWriter, message string, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(Response{
		Message: message,
		Status:  statusCode,
		Data:    data,
	})
}
