package model

type Cource struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Category    *Category `json:"category"`
}
