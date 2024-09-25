package lib

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
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

func ValidateSlug(fl validator.FieldLevel) bool {
	slug := fl.Field().String()
	re := regexp.MustCompile(`^[a-z0-9\-]+$`)
	return re.MatchString(slug)
}
