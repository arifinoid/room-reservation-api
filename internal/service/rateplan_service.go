package service

import (
	"github.com/arifinoid/room-reservation-api/internal/domain"
	"github.com/arifinoid/room-reservation-api/internal/repository"
)

type RatePlanService interface {
	CreateRateplan(rateplan domain.RatePlan) (int, error)
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
