package repository

import (
	"database/sql"

	"github.com/arifinoid/room-reservation-api/internal/domain"
)

type RatePlanRepository interface {
	Create(ratePlan domain.RatePlan) (int, error)
}

type ratePlanRepo struct {
	db *sql.DB
}

func NewRatePlanRepo(db *sql.DB) RatePlanRepository {
	return &ratePlanRepo{db: db}
}

func (r *ratePlanRepo) Create(ratePlan domain.RatePlan) (int, error) {
	var id int

	query := "INSERT INTO rateplans (room_id, name, slug, detail, price) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := r.db.QueryRow(query, ratePlan.RoomID, ratePlan.Name, ratePlan.Slug, ratePlan.Detail, ratePlan.Price).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
