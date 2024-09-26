package routes

import (
	"net/http"

	"github.com/arifinoid/room-reservation-api/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterCalendarRoutes(r *mux.Router, calendarHandler *handler.CalendarHandler) {
	subRouter := r.PathPrefix("/api/v1/calendars").Subrouter()
	subRouter.HandleFunc("", calendarHandler.CreateCalendar).Methods(http.MethodPost)
	subRouter.HandleFunc("", calendarHandler.GetCalendars).Methods(http.MethodGet)
	subRouter.HandleFunc("/{id}", calendarHandler.GetCalendarByID).Methods(http.MethodGet)
	subRouter.HandleFunc("/{id}", calendarHandler.UpdateCalendar).Methods(http.MethodPut)
	subRouter.HandleFunc("/{id}", calendarHandler.DeleteCalendar).Methods(http.MethodDelete)
}
