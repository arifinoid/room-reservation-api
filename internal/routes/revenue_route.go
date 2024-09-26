package routes

import (
	"net/http"

	"github.com/arifinoid/room-reservation-api/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterRevenueRoutes(r *mux.Router, revenueHandler *handler.RevenueHandler) {
	subRouter := r.PathPrefix("/api/v1/revenue").Subrouter()
	subRouter.HandleFunc("", revenueHandler.GetRevenue).Methods(http.MethodGet)
}
