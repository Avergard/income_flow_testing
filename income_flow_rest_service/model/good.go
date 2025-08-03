package model

type Good struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Count       int     `json:"count"`
	Weight      float64 `json:"weight"`
}
