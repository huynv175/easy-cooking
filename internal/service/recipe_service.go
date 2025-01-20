package service

import (
	"context"
	"easy-cooking/internal/models/dto"
	"easy-cooking/internal/repository"
	"gorm.io/gorm"
	"time"
)

type RecipeService interface {
	GetRecipes(ctx context.Context) ([]*dto.RecipeResponse, error)
	SearchRecipes(ctx context.Context, params dto.RecipeSearchRequest) (*dto.SearchRecipeResponse, error)
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

func (s *recipeService) GetRecipes(ctx context.Context) ([]*dto.RecipeResponse, error) {

	recipes, err := s.recipeRepository.GetRecipes(ctx)
	if err != nil {
		return nil, err
	}
	var result []*dto.RecipeResponse
	for _, recipe := range recipes {
		result = append(result, &dto.RecipeResponse{
			ID:           recipe.ID,
			Title:        recipe.Title,
			Description:  recipe.Description,
			Cuisine:      recipe.Cuisine,
			PhotoURL:     recipe.PhotoURL,
			CreatedAt:    recipe.CreatedAt,
			Ingredients:  nil,
			Instructions: nil,
		})
	}
	return result, nil
}

func (s *recipeService) SearchRecipes(ctx context.Context, params dto.RecipeSearchRequest) (*dto.SearchRecipeResponse, error) {
	// Validate request
	if err := params.Validate(); err != nil {
		return nil, err
	}

	searchCriteria := repository.SearchCriteria{
		Keyword:     params.Keyword,
		Ingredients: params.Ingredients,
		Cuisine:     params.Cuisine,
		Page:        params.Page,
		PageSize:    params.PageSize,
		SortBy:      params.SortBy,
		SortOrder:   params.SortOrder,
	}

	recipes, total, err := s.recipeRepository.SearchRecipes(ctx, searchCriteria)
	if err != nil {
		return nil, err
	}

	searchResponse := dto.NewSearchRecipeResponse(
		recipes,
		total,
		params.Page,
		params.PageSize,
	)

	return &searchResponse, nil
}
