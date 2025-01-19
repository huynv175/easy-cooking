package repository

import (
	"context"
	"easy-cooking/internal/models/do"
	"gorm.io/gorm"
)

type recipeRepository struct {
	db *gorm.DB
}

type RecipeRepository interface {
	GetRecipes(ctx context.Context) ([]*do.Recipe, error)
}

var _ RecipeRepository = (*recipeRepository)(nil)

func NewRecipeRepository(db *gorm.DB) *recipeRepository {
	return &recipeRepository{db: db}
}

func (r *recipeRepository) GetRecipes(ctx context.Context) ([]*do.Recipe, error) {
	var recipes []*do.Recipe
	if err := r.db.WithContext(ctx).Find(&recipes).Error; err != nil {
		return nil, err
	}
	return recipes, nil
}
