package handler

import (
	"encoding/json"
	"net/http"

	"github.com/arifinoid/room-reservation-api/internal/domain"
	"github.com/arifinoid/room-reservation-api/internal/lib"
	"github.com/arifinoid/room-reservation-api/internal/service"
	"github.com/go-playground/validator/v10"
)

type BookingHandler struct {
	BookingService service.BookingService
	validate       *validator.Validate
}

func NewBookingHandler(bookingService service.BookingService, validate *validator.Validate) *BookingHandler {
	return &BookingHandler{
		BookingService: bookingService,
		validate:       validate,
	}
}

func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking domain.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}

	if err := h.validate.Struct(booking); err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}

	id, err := h.BookingService.CreateBooking(booking)
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}
	lib.JSONResponse(w, struct {
		ID int `json:"id"`
	}{ID: id}, true, nil)
}

func (h *BookingHandler) GetBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := h.BookingService.GetBookings()
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}
	lib.JSONResponse(w, bookings, true, nil)
}
