package lib

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Data    T      `json:"data"`
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

func JSONResponse[T any](w http.ResponseWriter, data T, success bool, err error) {
	response := Response[T]{
		Data:    data,
		Success: success,
	}

	if err != nil {
		response.Error = err.Error()
	}
	w.Header().Set("Content-Type", "application/json")

	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(response)
}
