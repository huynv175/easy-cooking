package handler

import (
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
