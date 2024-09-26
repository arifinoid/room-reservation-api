package service

import (
	"github.com/arifinoid/room-reservation-api/internal/domain"
	"github.com/arifinoid/room-reservation-api/internal/repository"
)

type BookingService interface {
	CreateBooking(booking domain.Booking) (int, error)
	GetBookings(filter domain.BookingFilter) ([]domain.Booking, error)
	GetBookingByID(id int) (domain.Booking, error)
	UpdateBooking(id int, booking domain.Booking) error
	DeleteBooking(id int) error
}

type bookingService struct {
	bookingRepo repository.BookingRepository
}

func NewBookingService(bookingRepo repository.BookingRepository) BookingService {
	return &bookingService{
		bookingRepo: bookingRepo,
	}
}

func (s *bookingService) CreateBooking(booking domain.Booking) (int, error) {
	return s.bookingRepo.Create(booking)
}

func (s *bookingService) GetBookings(filter domain.BookingFilter) ([]domain.Booking, error) {
	return s.bookingRepo.GetAll(filter)
}

func (s *bookingService) GetBookingByID(id int) (domain.Booking, error) {
	return s.bookingRepo.GetByID(id)
}

func (s *bookingService) UpdateBooking(id int, booking domain.Booking) error {
	return s.bookingRepo.Update(id, booking)
}

func (s *bookingService) DeleteBooking(id int) error {
	return s.bookingRepo.Delete(id)
}
