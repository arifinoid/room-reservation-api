package service

import (
	"errors"

	"github.com/arifinoid/room-reservation-api/internal/domain"
	"github.com/arifinoid/room-reservation-api/internal/repository"
)

type RatePlanService interface {
	CreateRateplan(rateplan domain.RatePlan) (int, error)
	GetAllRateplans() ([]domain.RatePlan, error)
	GetRateplanByID(id int) (domain.RatePlan, error)
}

type ratePlanService struct {
	rateplanRepo repository.RatePlanRepository
}

func NewRatePlanService(rateplanRepo repository.RatePlanRepository) RatePlanService {
	return &ratePlanService{
		rateplanRepo: rateplanRepo,
	}
}

func (s *ratePlanService) CreateRateplan(rateplan domain.RatePlan) (int, error) {
	return s.rateplanRepo.Create(rateplan)
}

func (s *ratePlanService) GetAllRateplans() ([]domain.RatePlan, error) {
	return s.rateplanRepo.GetAll()
}

func (s *ratePlanService) GetRateplanByID(id int) (domain.RatePlan, error) {
	if id <= 0 {
		return domain.RatePlan{}, errors.New("invalid rateplan id")
	}
	return s.rateplanRepo.GetByID(id)
}
