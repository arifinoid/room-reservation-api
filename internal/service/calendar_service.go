package service

import (
	"github.com/arifinoid/room-reservation-api/internal/domain"
	"github.com/arifinoid/room-reservation-api/internal/repository"
)

type CalendarService interface {
	CreateCalendar(calendar domain.Calendar) (int, error)
}

type calendarService struct {
	calendarRepo repository.CalendarRepository
}

func NewCalendarService(calendarRepo repository.CalendarRepository) CalendarService {
	return &calendarService{
		calendarRepo: calendarRepo,
	}
}

func (s *calendarService) CreateCalendar(calendar domain.Calendar) (int, error) {
	return s.calendarRepo.Create(calendar)
}
