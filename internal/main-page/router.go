package mainPage

import (
	"github.com/erfanyousefi/codenight/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	service := NewService(config.GORMDB)
	handler := NewHandler(*service)
	mainGroup := r.Group("/")
	{
		mainGroup.GET("/", handler.GetInitData)
	}
}
