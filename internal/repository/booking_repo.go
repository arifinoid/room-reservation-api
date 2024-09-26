package repository

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/arifinoid/room-reservation-api/internal/domain"
)

type BookingRepository interface {
	Create(booking domain.Booking) (int, error)
	GetAll() ([]domain.Booking, error)
}

type bookingRepo struct {
	db *sql.DB
}

func NewBookingRepo(db *sql.DB) BookingRepository {
	return &bookingRepo{db: db}
}

func (r *bookingRepo) Create(booking domain.Booking) (int, error) {
	var id int

	now := time.Now().Format("20060102")
	randomPart := fmt.Sprintf("%06d", rand.Intn(1000000))
	reservationNumber := fmt.Sprintf("%s-%s", now, randomPart)

	query := `
	INSERT INTO bookings (room_id, rateplan_id, calendar_id, reservation_number, reservation_date, check_in, check_out, name, email, phone_number, country, total, payment_status)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	RETURNING id`
	err := r.db.QueryRow(query, booking.RoomID, booking.RateplanID, booking.CalendarID, reservationNumber, booking.ReservationDate, booking.CheckIn, booking.CheckOut,
		booking.Name, booking.Email, booking.PhoneNumber, booking.Country, booking.Total, booking.PaymentStatus).Scan(&id)
	if err != nil {
		return 0, err
	}

	reservationNumber = fmt.Sprintf("%s-%d", reservationNumber, id)
	updateQuery := `UPDATE bookings SET reservation_number = $1 WHERE id = $2`
	_, err = r.db.Exec(updateQuery, reservationNumber, id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *bookingRepo) GetAll() ([]domain.Booking, error) {
	var bookings []domain.Booking

	rows, err := r.db.Query("SELECT * FROM bookings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var booking domain.Booking
		if err := rows.Scan(&booking.ID, &booking.RoomID, &booking.RateplanID, &booking.CalendarID, &booking.ReservationNumber, &booking.ReservationDate, &booking.CheckIn, &booking.CheckOut,
			&booking.Name, &booking.Email, &booking.PhoneNumber, &booking.Country, &booking.Total, &booking.PaymentStatus); err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}
