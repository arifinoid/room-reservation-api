package domain

type Room struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Description  string `json:"description"`
	Feature      string `json:"feature"`
	Published    bool   `json:"published"`
	Availability int    `json:"availability"`
	Images       string `json:"images"`
}
