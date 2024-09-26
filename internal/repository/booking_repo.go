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
	GetByID(id int) (domain.Booking, error)
	Update(id int, booking domain.Booking) error
	Delete(id int) error
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

	updateAvailabilityQuery := `
	UPDATE calendars
	SET availability = availability - 1
	WHERE room_id = $1
	  AND rateplan_id = $2
	  AND date BETWEEN $3 AND $4`
	_, err = r.db.Exec(updateAvailabilityQuery, booking.RoomID, booking.RateplanID, booking.CheckIn, booking.CheckOut)
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

func (r *bookingRepo) GetByID(id int) (domain.Booking, error) {
	var booking domain.Booking
	row := r.db.QueryRow("SELECT * FROM bookings WHERE id = $1", id)
	if err := row.Scan(&booking.ID, &booking.RoomID, &booking.RateplanID, &booking.CalendarID, &booking.ReservationNumber, &booking.ReservationDate, &booking.CheckIn, &booking.CheckOut,
		&booking.Name, &booking.Email, &booking.PhoneNumber, &booking.Country, &booking.Total, &booking.PaymentStatus); err != nil {
		return booking, err
	}
	return booking, nil
}

func (r *bookingRepo) Update(id int, booking domain.Booking) error {
	query := `UPDATE bookings SET name = $1, email = $2, phone_number = $3, country = $4, total = $5, payment_status = $6 WHERE id = $7`
	_, err := r.db.Exec(query, booking.Name, booking.Email, booking.PhoneNumber, booking.Country, booking.Total, booking.PaymentStatus, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *bookingRepo) Delete(id int) error {
	var roomID, rateplanID int
	var checkIn, checkOut time.Time

	query := `SELECT room_id, rateplan_id, check_in, check_out FROM bookings WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&roomID, &rateplanID, &checkIn, &checkOut)
	if err != nil {
		return err
	}

	updateAvailabilityQuery := `
	UPDATE calendars
	SET availability = availability + 1
	WHERE room_id = $1
	  AND rateplan_id = $2
	  AND date BETWEEN $3 AND $4`
	_, err = r.db.Exec(updateAvailabilityQuery, roomID, rateplanID, checkIn, checkOut)
	if err != nil {
		return err
	}

	deleteQuery := `DELETE FROM bookings WHERE id = $1`
	_, err = r.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	return nil
}
