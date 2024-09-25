package service

import (
	"errors"

	"github.com/arifinoid/room-reservation-api/internal/domain"
	"github.com/arifinoid/room-reservation-api/internal/repository"
)

type RoomService interface {
	GetAllRooms() ([]domain.Room, error)
	GetRoomByID(id int) (domain.Room, error)
	CreateRoom(room domain.Room) (int, error)
	UpdateRoom(id int, room domain.Room) error
	DeleteRoom(id int) error
}

type roomService struct {
	roomRepo repository.RoomRepository
}

func NewRoomService(roomRepo repository.RoomRepository) RoomService {
	return &roomService{
		roomRepo: roomRepo,
	}
}

func (s *roomService) GetAllRooms() ([]domain.Room, error) {
	return s.roomRepo.GetAll()
}

func (s *roomService) GetRoomByID(id int) (domain.Room, error) {
	if id <= 0 {
		return domain.Room{}, errors.New("invalid room id")
	}
	return s.roomRepo.GetByID(id)
}

func (s *roomService) CreateRoom(room domain.Room) (int, error) {
	if room.Name == "" {
		return 0, errors.New("invalid room name")
	}
	return s.roomRepo.Create(room)
}

func (s *roomService) UpdateRoom(id int, room domain.Room) error {
	if id <= 0 {
		return errors.New("invalid room id")
	}
	return s.roomRepo.Update(id, room)
}

func (s *roomService) DeleteRoom(id int) error {
	if id <= 0 {
		return errors.New("invalid room id")
	}
	return s.roomRepo.Delete(id)
}
