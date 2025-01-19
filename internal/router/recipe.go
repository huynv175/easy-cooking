package router

import (
	"easy-cooking/internal/handler"
	"github.com/gin-gonic/gin"
)

func RecipeRouter(r *gin.Engine, handler *handler.Handler) {
	productGroup := r.Group("v0/recipes")
	{
		productGroup.GET("", handler.GetRecipes)
	}
}
