package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/arifinoid/room-reservation-api/internal/domain"
	"github.com/arifinoid/room-reservation-api/internal/lib"
	"github.com/arifinoid/room-reservation-api/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type RatePlanHandler struct {
	RatePlanService service.RatePlanService
	validate        *validator.Validate
}

func NewRatePlanHandler(rateplanService service.RatePlanService, validate *validator.Validate) *RatePlanHandler {
	return &RatePlanHandler{
		RatePlanService: rateplanService,
		validate:        validate,
	}
}

func (h *RatePlanHandler) CreateRateplan(w http.ResponseWriter, r *http.Request) {
	var rateplan domain.RatePlan
	if err := json.NewDecoder(r.Body).Decode(&rateplan); err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}

	if err := h.validate.Struct(rateplan); err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}

	id, err := h.RatePlanService.CreateRateplan(rateplan)
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}
	lib.JSONResponse(w, id, true, nil)
}

func (h *RatePlanHandler) GetRateplanByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, errors.New("invalid rateplan id"))
		return
	}
	rateplan, err := h.RatePlanService.GetRateplanByID(id)
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}
	lib.JSONResponse(w, rateplan, true, nil)
}

func (h *RatePlanHandler) GetRateplans(w http.ResponseWriter, r *http.Request) {
	rateplans, err := h.RatePlanService.GetAllRateplans()
	if err != nil {
		lib.JSONResponse(w, struct{}{}, false, err)
		return
	}
	lib.JSONResponse(w, rateplans, true, nil)
}
