package routes

import (
	"net/http"

	"github.com/arifinoid/room-reservation-api/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterBookingRoutes(r *mux.Router, bookingHandler *handler.BookingHandler) {
	subRouter := r.PathPrefix("/api/v1/bookings").Subrouter()
	subRouter.HandleFunc("", bookingHandler.CreateBooking).Methods(http.MethodPost)
	subRouter.HandleFunc("", bookingHandler.GetBookings).Methods(http.MethodGet)
}
