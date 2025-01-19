package handler

import (
	"easy-cooking/internal/models/dto"
	"easy-cooking/internal/service"
	"time"

	"gorm.io/gorm"
)

type Handler struct {
	DB            *gorm.DB
	recipeService service.RecipeService
}

func NewHandler(db *gorm.DB, jwtConfig dto.JWTConfig, timeout time.Duration) *Handler {
	return &Handler{
		DB:            db,
		recipeService: service.NewRecipeService(db, jwtConfig, timeout),
	}
}
