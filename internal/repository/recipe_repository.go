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
	SearchRecipes(ctx context.Context, criteria SearchCriteria) ([]*do.Recipe, int64, error)
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

type SearchCriteria struct {
	Keyword     string
	Ingredients []string
	Cuisine     string
	Page        int
	PageSize    int
	SortBy      string
	SortOrder   string
}

func (r *recipeRepository) SearchRecipes(ctx context.Context, criteria SearchCriteria) ([]*do.Recipe, int64, error) {
	query := r.db.Model(&do.Recipe{})

	if criteria.Keyword != "" {
		query = query.Where("title ILIKE ? OR description ILIKE ?",
			"%"+criteria.Keyword+"%",
			"%"+criteria.Keyword+"%")
	}

	if criteria.Cuisine != "" {
		query = query.Where("cuisine = ?", criteria.Cuisine)
	}

	if len(criteria.Ingredients) > 0 {
		query = query.Joins("JOIN recipe_ingredients ri ON ri.recipe_id = recipes.id").
			Joins("JOIN ingredients i ON i.id = ri.ingredient_id").
			Where("i.name IN ?", criteria.Ingredients).
			Group("recipes.id")
	}

	if criteria.SortBy != "" {
		orderClause := criteria.SortBy
		if criteria.SortOrder != "" {
			orderClause += " " + criteria.SortOrder
		}
		query = query.Order(orderClause)
	}

	var total int64
	countQuery := query
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (criteria.Page - 1) * criteria.PageSize
	query = query.Offset(offset).Limit(criteria.PageSize)

	var recipes []*do.Recipe
	if err := query.Find(&recipes).Error; err != nil {
		return nil, 0, err
	}

	return recipes, total, nil
}
