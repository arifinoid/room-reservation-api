package repository

import (
	"database/sql"

	"github.com/arifinoid/room-reservation-api/internal/domain"
)

type RatePlanRepository interface {
	Create(ratePlan domain.RatePlan) (int, error)
	GetAll() ([]domain.RatePlan, error)
	GetByID(id int) (domain.RatePlan, error)
	Update(id int, ratePlan domain.RatePlan) error
	Delete(id int) error
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

func (r *ratePlanRepo) GetAll() ([]domain.RatePlan, error) {
	var ratePlans []domain.RatePlan
	query := `SELECT id, room_id, name, slug, detail, price FROM rateplans`
	rows, err := r.db.Query(query)
	if err != nil {
		return ratePlans, err
	}
	defer rows.Close()
	for rows.Next() {
		var ratePlan domain.RatePlan
		if err := rows.Scan(&ratePlan.ID, &ratePlan.RoomID, &ratePlan.Name, &ratePlan.Slug, &ratePlan.Detail, &ratePlan.Price); err != nil {
			return ratePlans, err
		}
		ratePlans = append(ratePlans, ratePlan)
	}
	return ratePlans, nil
}

func (r *ratePlanRepo) GetByID(id int) (domain.RatePlan, error) {
	var ratePlan domain.RatePlan
	query := `SELECT id, room_id, name, slug, detail, price FROM rateplans WHERE id = $1`
	if err := r.db.QueryRow(query, id).Scan(&ratePlan.ID, &ratePlan.RoomID, &ratePlan.Name, &ratePlan.Slug, &ratePlan.Detail, &ratePlan.Price); err != nil {
		return ratePlan, err
	}
	return ratePlan, nil
}

func (r *ratePlanRepo) Update(id int, ratePlan domain.RatePlan) error {
	query := "UPDATE rateplans SET room_id = $1, name = $2, slug = $3, detail = $4, price = $5 WHERE id = $6"
	_, err := r.db.Exec(query, ratePlan.RoomID, ratePlan.Name, ratePlan.Slug, ratePlan.Detail, ratePlan.Price, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ratePlanRepo) Delete(id int) error {
	query := "DELETE FROM rateplans WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
