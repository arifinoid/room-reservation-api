package domain

type Calendar struct {
	ID           int     `json:"id"`
	RoomID       int     `json:"room_id" validate:"required"`
	RatePlanID   int     `json:"rateplan_id" validate:"required"`
	Date         string  `json:"date" validate:"required"`
	Availability int     `json:"availability"`
	Price        float64 `json:"price"`
}
