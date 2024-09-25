package routes

import (
	"net/http"

	"github.com/arifinoid/room-reservation-api/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterRoomRoutes(r *mux.Router, roomHandler *handler.RoomHandler) {
	subRouter := r.PathPrefix("/api/v1/rooms").Subrouter()
	subRouter.HandleFunc("", roomHandler.GetRooms).Methods(http.MethodGet)
}
