package service

import (
	"github.com/arifinoid/room-reservation-api/internal/domain"
	"github.com/arifinoid/room-reservation-api/internal/repository"
)

type BookingService interface {
	CreateBooking(booking domain.Booking) (int, error)
	GetBookings() ([]domain.Booking, error)
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

func (s *bookingService) GetBookings() ([]domain.Booking, error) {
	return s.bookingRepo.GetAll()
}
