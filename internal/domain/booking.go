package domain

import "time"

type Booking struct {
	ID                int       `json:"id"`
	RoomID            int       `json:"room_id" validate:"required"`
	RateplanID        int       `json:"rateplan_id" validate:"required"`
	CalendarID        int       `json:"calendar_id" validate:"required"`
	ReservationNumber string    `json:"reservation_number"`
	ReservationDate   time.Time `json:"reservation_date" validate:"required"`
	CheckIn           time.Time `json:"check_in" validate:"required"`
	CheckOut          time.Time `json:"check_out" validate:"required,gtfield=CheckIn"`
	Name              string    `json:"name" validate:"required,min=2,max=100"`
	Email             string    `json:"email" validate:"required,email"`
	PhoneNumber       string    `json:"phone_number" validate:"required,e164"`
	Country           string    `json:"country" validate:"required"`
	Total             float64   `json:"total" validate:"required,gte=0"`
	PaymentStatus     string    `json:"payment_status" validate:"required,oneof='paid' 'pending' 'canceled'"`
}

type BookingFilter struct {
	ReservationDateFrom   string `json:"reservation_date_from"`
	ReservationDateTo     string `json:"reservation_date_to"`
	CheckInDate           string `json:"check_in_date"`
	CheckOutDate          string `json:"check_out_date"`
	GuestName             string `json:"guest_name"`
	GuestCountry          string `json:"guest_country"`
	PaymentStatus         string `json:"payment_status"`
	ReservationNumberFrom string `json:"reservation_number_from"`
	ReservationNumberTo   string `json:"reservation_number_to"`
}
