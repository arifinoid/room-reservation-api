package repository

import (
	"database/sql"
	"errors"

	"github.com/arifinoid/room-reservation-api/internal/domain"
)

type CalendarRepository interface {
	Create(calendar domain.Calendar) (int, error)
	GetAll() ([]domain.Calendar, error)
	GetByID(id int) (domain.Calendar, error)
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

func (r *calendarRepo) GetAll() ([]domain.Calendar, error) {
	var calendars []domain.Calendar
	query := `SELECT id, room_id, rateplan_id, date, availability, price FROM calendars`
	rows, err := r.db.Query(query)
	if err != nil {
		return calendars, err
	}
	defer rows.Close()
	for rows.Next() {
		var calendar domain.Calendar
		if err := rows.Scan(&calendar.ID, &calendar.RoomID, &calendar.RatePlanID, &calendar.Date, &calendar.Availability, &calendar.Price); err != nil {
			return calendars, err
		}
		calendars = append(calendars, calendar)
	}
	return calendars, nil
}

func (r *calendarRepo) GetByID(id int) (domain.Calendar, error) {
	var calendar domain.Calendar
	query := `SELECT id, room_id, rateplan_id, date, availability, price FROM calendars WHERE id = $1`
	if err := r.db.QueryRow(query, id).Scan(&calendar.ID, &calendar.RoomID, &calendar.RatePlanID, &calendar.Date, &calendar.Availability, &calendar.Price); err != nil {
		return calendar, err
	}
	return calendar, nil
}
