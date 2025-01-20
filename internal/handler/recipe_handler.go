package handler

import (
	"easy-cooking/internal/models/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetRecipes(c *gin.Context) {
	recipes, err := h.recipeService.GetRecipes(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recipes"})
		return
	}
	c.JSON(http.StatusOK, recipes)
}

func (h *Handler) GetRecipe(c *gin.Context) {
	recipes, err := h.recipeService.GetRecipes(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recipes"})
		return
	}
	c.JSON(http.StatusOK, recipes)
}

func (h *Handler) SearchRecipes(c *gin.Context) {
	var searchParams dto.RecipeSearchRequest

	if err := c.ShouldBind(&searchParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid search parameters",
			"details": err.Error(),
		})
		return
	}

	searchResponse, err := h.recipeService.SearchRecipes(c.Request.Context(), searchParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch recipes",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, searchResponse)
}
