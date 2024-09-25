package repository

import (
	"database/sql"
	"errors"

	"github.com/arifinoid/room-reservation-api/internal/domain"
)

type CalendarRepository interface {
	Create(calendar domain.Calendar) (int, error)
}

type calendarRepo struct {
	db *sql.DB
}

func NewCalendarRepo(db *sql.DB) CalendarRepository {
	return &calendarRepo{db: db}
}

func (r *calendarRepo) Create(calendar domain.Calendar) (int, error) {
	var id int

	var roomAvailability int
	if err := r.db.QueryRow("SELECT availability FROM rooms WHERE id = $1", calendar.RoomID).Scan(&roomAvailability); err != nil {
		return 0, err
	}

	var ratePlanPrice float64
	if err := r.db.QueryRow("SELECT price FROM rateplans WHERE id = $1", calendar.RatePlanID).Scan(&ratePlanPrice); err != nil {
		return 0, err
	}

	if calendar.Availability == 0 {
		calendar.Availability = roomAvailability
	} else if calendar.Availability > roomAvailability {
		return 0, errors.New("calendar availability cannot be greater than room availability")
	}

	if calendar.Price == 0 {
		calendar.Price = ratePlanPrice
	} else if calendar.Price > ratePlanPrice {
		return 0, errors.New("calendar price cannot be greater than rate plan price")
	}

	query := "INSERT INTO calendars (room_id, rateplan_id, date, availability, price) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := r.db.QueryRow(query, calendar.RoomID, calendar.RatePlanID, calendar.Date, calendar.Availability, calendar.Price).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
