package domain

type RatePlan struct {
	ID     int     `json:"id"`
	RoomID int     `json:"room_id" validate:"required"`
	Name   string  `json:"name" validate:"required"`
	Slug   string  `json:"slug" validate:"required,slug"`
	Detail string  `json:"detail"`
	Price  float64 `json:"price" validate:"required"`
}
