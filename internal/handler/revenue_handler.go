package handler

import (
	"net/http"

	"github.com/arifinoid/room-reservation-api/internal/lib"
	"github.com/arifinoid/room-reservation-api/internal/service"
)

type RevenueHandler struct {
	RevenueService service.RevenueService
}

func NewRevenueHandler(revenueService service.RevenueService) *RevenueHandler {
	return &RevenueHandler{
		RevenueService: revenueService,
	}
}

func (h *RevenueHandler) GetRevenue(w http.ResponseWriter, r *http.Request) {
	revenue, err := h.RevenueService.GetRevenue()
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}
	lib.JSONResponse(w, revenue, true, nil)
}
