package auth

import (
	"github.com/erfanyousefi/codenight/config"
	"github.com/erfanyousefi/codenight/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, jwtSecret string) {
	authService := AuthService{
		DB:        config.GORMDB,
		JWTSecret: jwtSecret,
	}

	authGroup := r.Group("/auth")
	{
		authGroup.GET("/profile", middleware.AuthMiddleware(jwtSecret), ProfileHandler())
		authGroup.POST("/register", RegisterHandler(&authService))
		authGroup.POST("/login", LoginHandler(&authService))
	}
}
