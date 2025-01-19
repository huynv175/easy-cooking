package service

import (
	"context"
	"easy-cooking/internal/models/dto"
	"easy-cooking/internal/repository"
	"gorm.io/gorm"
	"time"
)

type RecipeService interface {
	GetRecipes(ctx context.Context) ([]*dto.Recipe, error)
}

type recipeService struct {
	timeout          time.Duration
	recipeRepository repository.RecipeRepository
	jwtConfig        dto.JWTConfig
}

func NewRecipeService(db *gorm.DB, jwtConfig dto.JWTConfig, timeout time.Duration) *recipeService {
	return &recipeService{
		timeout:          timeout,
		recipeRepository: repository.NewRecipeRepository(db),
		jwtConfig:        jwtConfig,
	}
}

var _ RecipeService = (*recipeService)(nil)

func (s *recipeService) GetRecipes(ctx context.Context) ([]*dto.Recipe, error) {

	recipes, err := s.recipeRepository.GetRecipes(ctx)
	if err != nil {
		return nil, err
	}
	var result []*dto.Recipe
	for _, recipe := range recipes {
		result = append(result, &dto.Recipe{
			ID:          recipe.ID,
			Title:       "",
			Description: recipe.Description,
			Cuisine:     "",
			PhotoURL:    "",
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
			DeletedAt:   nil,
		})
	}
	return result, nil
}
