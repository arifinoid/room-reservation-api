package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/arifinoid/room-reservation-api/internal/domain"
	"github.com/arifinoid/room-reservation-api/internal/lib"
	"github.com/arifinoid/room-reservation-api/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
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

func (h *BookingHandler) GetBookingByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, errors.New("invalid booking id"))
		return
	}

	booking, err := h.BookingService.GetBookingByID(id)
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}

	lib.JSONResponse(w, booking, true, nil)
}

func (h *BookingHandler) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, errors.New("invalid booking id"))
		return
	}

	var booking domain.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}

	if err := h.validate.Struct(booking); err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}

	err = h.BookingService.UpdateBooking(id, booking)
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}
	lib.JSONResponse(w, struct{}{}, true, nil)
}

func (h *BookingHandler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, errors.New("invalid booking id"))
		return
	}

	err = h.BookingService.DeleteBooking(id)
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}
	lib.JSONResponse(w, struct{}{}, true, nil)
}
