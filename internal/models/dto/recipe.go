package dto

import "time"

type Recipe struct {
	ID          uint64     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Cuisine     string     `json:"cuisine"`
	PhotoURL    string     `json:"photo_url"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	//Ingredients  []Ingredient  `json:"ingredients,omitempty"`
	//Instructions []Instruction `json:"instructions,omitempty"`
}
