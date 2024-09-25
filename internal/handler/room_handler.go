package handler

import (
	"net/http"

	"github.com/arifinoid/room-reservation-api/internal/lib"
	"github.com/arifinoid/room-reservation-api/internal/service"
)

type RoomHandler struct {
	RoomService service.RoomService
}

func NewRoomHandler(roomService service.RoomService) *RoomHandler {
	return &RoomHandler{
		RoomService: roomService,
	}
}

func (h *RoomHandler) GetRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := h.RoomService.GetAll()
	if err != nil {

		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}

	lib.JSONResponse(w, rooms, true, nil)
}
