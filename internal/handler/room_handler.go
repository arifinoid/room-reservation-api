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

type RoomHandler struct {
	RoomService service.RoomService
	validate    *validator.Validate
}

func NewRoomHandler(roomService service.RoomService, validate *validator.Validate) *RoomHandler {
	return &RoomHandler{
		RoomService: roomService,
		validate:    validate,
	}
}

func (h *RoomHandler) GetRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := h.RoomService.GetAllRooms()
	if err != nil {

		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}

	lib.JSONResponse(w, rooms, true, nil)
}

func (h *RoomHandler) GetRoom(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, errors.New("invalid room id"))
		return
	}

	room, err := h.RoomService.GetRoomByID(id)
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}
	lib.JSONResponse(w, room, true, nil)
}

func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var room domain.Room
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}

	if err := h.validate.Struct(room); err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}

	id, err := h.RoomService.CreateRoom(room)
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}
	lib.JSONResponse(w, struct {
		ID int `json:"id"`
	}{ID: id}, true, nil)
}
