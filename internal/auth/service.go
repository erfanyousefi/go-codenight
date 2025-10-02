package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/erfanyousefi/codenight/internal/user"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB        *gorm.DB
	JWTSecret string
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func (s *AuthService) Register(firstname, lastname, email, password string) error {
	var existUser user.User
	if err := s.DB.Where("email = ?", email).First(&existUser).Error; err == nil {
		return errors.New("email already exist")
	} else if err != gorm.ErrRecordNotFound {
		return err
	}
	fmt.Println(existUser)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u := user.User{
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		Password:  string(hashedPassword),
		Balance:   10000,
	}
	return s.DB.Create(&u).Error
}

func (s *AuthService) Login(email, password string) (string, error) {
	var user user.User

	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	claim := &Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
