package routes

import (
	"net/http"

	"github.com/arifinoid/room-reservation-api/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterRoomRoutes(r *mux.Router, roomHandler *handler.RoomHandler) {
	subRouter := r.PathPrefix("/api/v1/rooms").Subrouter()
	subRouter.HandleFunc("", roomHandler.GetRooms).Methods(http.MethodGet)
	subRouter.HandleFunc("/{id}", roomHandler.GetRoom).Methods(http.MethodGet)
	subRouter.HandleFunc("", roomHandler.CreateRoom).Methods(http.MethodPost)
	subRouter.HandleFunc("/{id}", roomHandler.UpdateRoom).Methods(http.MethodPut)
	subRouter.HandleFunc("/{id}", roomHandler.DeleteRoom).Methods(http.MethodDelete)
}
