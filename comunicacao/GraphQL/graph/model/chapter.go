package model

type Chapter struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Cource   *Cource   `json:"cource"`
	Category *Category `json:"category"`
}
