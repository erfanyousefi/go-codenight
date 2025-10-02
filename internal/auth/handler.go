package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Firstname string `json:"firstname" binding:"required,min=3,max=50"`
	Lastname  string `json:"lastname" binding:"required,min=3,max=50"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6,max=25"`
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=25"`
}

func RegisterHandler(s *AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req RegisterRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := s.Register(req.Firstname, req.Lastname, req.Email, req.Password); err != nil {
			if err.Error() == "email already exist" {
				ctx.JSON(http.StatusConflict, gin.H{"error": "email already exist"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
	}
}

func LoginHandler(s *AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req LoginRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		token, err := s.Login(req.Email, req.Password)
		if err != nil {
			if err.Error() == "invalid credentials" {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "email or password is invalid"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func ProfileHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if userId, ok := ctx.Get("userId"); ok {
			ctx.JSON(http.StatusOK, gin.H{"userId": userId})
			return
		}
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "login..."})
	}
}
