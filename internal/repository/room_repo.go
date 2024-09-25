package repository

import (
	"database/sql"

	"github.com/arifinoid/room-reservation-api/internal/domain"
)

type RoomRepository interface {
	GetAll() ([]domain.Room, error)
}

type roomRepo struct {
	db *sql.DB
}

func NewRoomRepo(db *sql.DB) RoomRepository {
	return &roomRepo{db: db}
}

func (r *roomRepo) GetAll() ([]domain.Room, error) {
	var rooms []domain.Room
	rows, err := r.db.Query("SELECT * FROM rooms")
	if err != nil {
		return rooms, err
	}
	defer rows.Close()
	for rows.Next() {
		var room domain.Room
		if err := rows.Scan(&room.ID, &room.Name, &room.Slug, &room.Description, &room.Feature, &room.Published, &room.Availability, &room.Images); err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}
