package routes

import (
	"net/http"

	"github.com/arifinoid/room-reservation-api/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterRatePlanRoutes(r *mux.Router, rateplanHandler *handler.RatePlanHandler) {
	subRouter := r.PathPrefix("/api/v1/rateplans").Subrouter()
	subRouter.HandleFunc("", rateplanHandler.CreateRateplan).Methods(http.MethodPost)
	subRouter.HandleFunc("", rateplanHandler.GetRateplans).Methods(http.MethodGet)
	subRouter.HandleFunc("/{id}", rateplanHandler.GetRateplanByID).Methods(http.MethodGet)
}
