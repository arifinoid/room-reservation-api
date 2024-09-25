package domain

type Room struct {
	ID           int                    `json:"id"`
	Name         string                 `json:"name" validate:"required,min=3,max=100"`
	Slug         string                 `json:"slug" validate:"required,slug"`
	Description  string                 `json:"description" validate:"required,min=3,max=200"`
	Feature      map[string]interface{} `json:"feature" validate:"required"`
	Published    bool                   `json:"published" validate:"-"`
	Availability int                    `json:"availability" validate:"required,min=1"`
	Images       []string               `json:"images,omitempty" validate:"dive,uri"`
}
