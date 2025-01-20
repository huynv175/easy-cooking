package dto

import "easy-cooking/internal/models/do"

type IngredientResponse struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
	Unit     string `json:"unit"`
}

func ToIngredientResponse(ing do.Ingredient) IngredientResponse {
	return IngredientResponse{
		ID:   ing.ID,
		Name: ing.Name,
	}
}
