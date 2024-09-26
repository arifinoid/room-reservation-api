package service

import (
	"github.com/arifinoid/room-reservation-api/internal/domain"
	"github.com/arifinoid/room-reservation-api/internal/repository"
)

type RevenueService interface {
	GetRevenue() (domain.Revenue, error)
}

type revenueService struct {
	revenueRepo repository.RevenueRepository
}

func NewRevenueService(revenueRepo repository.RevenueRepository) RevenueService {
	return &revenueService{
		revenueRepo: revenueRepo,
	}
}

func (s *revenueService) GetRevenue() (domain.Revenue, error) {
	return s.revenueRepo.GetRevenue()
}
