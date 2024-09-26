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

type CalendarHandler struct {
	CalendarService service.CalendarService
	validate        *validator.Validate
}

func NewCandendarHandler(calendarService service.CalendarService, validate *validator.Validate) *CalendarHandler {
	return &CalendarHandler{
		CalendarService: calendarService,
		validate:        validate,
	}
}

func (h *CalendarHandler) CreateCalendar(w http.ResponseWriter, r *http.Request) {
	var calendar domain.Calendar
	if err := json.NewDecoder(r.Body).Decode(&calendar); err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}

	if err := h.validate.Struct(calendar); err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}

	id, err := h.CalendarService.CreateCalendar(calendar)
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}

	lib.JSONResponse(w, struct {
		ID int `json:"id"`
	}{
		ID: id,
	}, true, nil)
}

func (h *CalendarHandler) GetCalendars(w http.ResponseWriter, r *http.Request) {
	cals, err := h.CalendarService.GetCalendars()
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}
	lib.JSONResponse(w, cals, true, nil)
}

func (h *CalendarHandler) GetCalendarByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, errors.New("invalid calendar id"))
		return
	}
	cal, err := h.CalendarService.GetCalendar(id)
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}
	lib.JSONResponse(w, cal, true, nil)
}

func (h *CalendarHandler) UpdateCalendar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, errors.New("invalid calendar id"))
		return
	}

	var calendar domain.Calendar
	if err := json.NewDecoder(r.Body).Decode(&calendar); err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}

	if err := h.validate.Struct(calendar); err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}

	calendar.ID = id

	err = h.CalendarService.UpdateCalendar(id, calendar)
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}
	lib.JSONResponse(w, struct{}{}, true, nil)
}

func (h *CalendarHandler) DeleteCalendar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, errors.New("invalid calendar id"))
		return
	}
	err = h.CalendarService.DeleteCalendar(id)
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}
	lib.JSONResponse(w, struct{}{}, true, nil)
}
