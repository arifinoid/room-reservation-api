package repository

import (
	"database/sql"
	"encoding/json"

	"github.com/arifinoid/room-reservation-api/internal/domain"
	"github.com/lib/pq"
)

type RoomRepository interface {
	GetAll() ([]domain.Room, error)
	GetByID(id int) (domain.Room, error)
	Create(room domain.Room) (int, error)
	Update(id int, room domain.Room) error
	Delete(id int) error
}

type roomRepo struct {
	db *sql.DB
}

func NewRoomRepo(db *sql.DB) RoomRepository {
	return &roomRepo{db: db}
}

func (r *roomRepo) GetAll() ([]domain.Room, error) {
	var rooms []domain.Room
	query := `SELECT id, name, slug, description, feature, published, availability, images FROM rooms`
	rows, err := r.db.Query(query)
	if err != nil {
		return rooms, err
	}
	defer rows.Close()
	for rows.Next() {
		var room domain.Room
		var featureJSON string

		if err := rows.Scan(&room.ID, &room.Name, &room.Slug, &room.Description, &featureJSON, &room.Published, &room.Availability, pq.Array(&room.Images)); err != nil {
			return rooms, err
		}

		if err := json.Unmarshal([]byte(featureJSON), &room.Feature); err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}

func (r *roomRepo) GetByID(id int) (domain.Room, error) {
	var room domain.Room
	var featureJSON string

	query := `SELECT id, name, slug, description, feature, published, availability, images FROM rooms WHERE id = $1`

	if err := r.db.QueryRow(query, id).Scan(&room.ID, &room.Name, &room.Slug, &room.Description, &featureJSON, &room.Published, &room.Availability, pq.Array(&room.Images)); err != nil {
		return room, err
	}

	err := json.Unmarshal([]byte(featureJSON), &room.Feature)
	if err != nil {
		return room, err
	}
	return room, nil
}

func (r *roomRepo) Create(room domain.Room) (int, error) {
	var id int

	featureJSON, err := json.Marshal(room.Feature)
	if err != nil {
		return 0, err
	}

	var imagesArray interface{}
	if len(room.Images) == 0 {
		imagesArray = pq.Array(nil)
	} else {
		imagesArray = pq.Array(room.Images)
	}

	query := "INSERT INTO rooms (name, slug, description, feature, published, availability, images) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	err = r.db.QueryRow(query, room.Name, room.Slug, room.Description, string(featureJSON), room.Published, room.Availability, imagesArray).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *roomRepo) Update(id int, room domain.Room) error {
	featureJSON, err := json.Marshal(room.Feature)
	if err != nil {
		return err
	}

	var imagesArray interface{}
	if len(room.Images) == 0 {
		imagesArray = pq.Array(nil)
	} else {
		imagesArray = pq.Array(room.Images)
	}

	query := "UPDATE rooms SET name = $1, slug = $2, description = $3, feature = $4, published = $5, availability = $6, images = $7 WHERE id = $8"
	_, err = r.db.Exec(query, room.Name, room.Slug, room.Description, string(featureJSON), room.Published, room.Availability, imagesArray, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *roomRepo) Delete(id int) error {
	query := "DELETE FROM rooms WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
