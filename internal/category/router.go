package category

import (
	"github.com/erfanyousefi/codenight/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	service := NewService(config.GORMDB)
	handler := NewHandler(*service)

	categoryGroup := r.Group("/category")
	{
		// categoryGroup.Use(middleware.AuthMiddleware(*config.JwtSecretKey))
		categoryGroup.POST("/", handler.CreateCategory)
		categoryGroup.GET("/", handler.GetAllCategories)
		categoryGroup.GET("/parents", handler.GetAllParents)
		categoryGroup.GET("/:id", handler.GetByID)
		categoryGroup.PUT("/:id", handler.UpdateCategory)
		categoryGroup.DELETE("/:id", handler.DeleteCategory)
	}

}
