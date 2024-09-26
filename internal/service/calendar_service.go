package service

import (
	"errors"

	"github.com/arifinoid/room-reservation-api/internal/domain"
	"github.com/arifinoid/room-reservation-api/internal/repository"
)

type CalendarService interface {
	CreateCalendar(calendar domain.Calendar) (int, error)
	GetCalendars() ([]domain.Calendar, error)
	GetCalendar(id int) (domain.Calendar, error)
	UpdateCalendar(id int, calendar domain.Calendar) error
	DeleteCalendar(id int) error
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

func (s *calendarService) GetCalendars() ([]domain.Calendar, error) {
	return s.calendarRepo.GetAll()
}

func (s *calendarService) GetCalendar(id int) (domain.Calendar, error) {
	if id <= 0 {
		return domain.Calendar{}, errors.New("invalid id")
	}
	return s.calendarRepo.GetByID(id)
}

func (s *calendarService) UpdateCalendar(id int, calendar domain.Calendar) error {
	if id <= 0 {
		return errors.New("invalid id")
	}
	return s.calendarRepo.Update(id, calendar)
}

func (s *calendarService) DeleteCalendar(id int) error {
	if id <= 0 {
		return errors.New("invalid id")
	}
	return s.calendarRepo.Delete(id)
}
