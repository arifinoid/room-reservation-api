package handler

import (
	"encoding/json"
	"net/http"

	"github.com/arifinoid/room-reservation-api/internal/domain"
	"github.com/arifinoid/room-reservation-api/internal/lib"
	"github.com/arifinoid/room-reservation-api/internal/service"
	"github.com/go-playground/validator/v10"
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
