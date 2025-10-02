package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ConfigEnv struct {
	ServerPort string
	DBUser     string
	DBPass     string
	DBHost     string
	DBPort     string
	DBName     string
	JWTSecret  string
}

var JwtSecretKey *string

func LoadConfig() ConfigEnv {
	if err := godotenv.Load(); err != nil {
		fmt.Println("cannot load environment variable ")
	}
	cfg := ConfigEnv{
		DBUser:     getEnv("DB_USER", "root"),
		DBPass:     getEnv("DB_PASS", "root"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBName:     getEnv("DB_NAME", "codenight"),
		ServerPort: getEnv("PORT", "8080"),
		JWTSecret:  getEnv("JWT_SECRET", "secretvaluekey*"),
	}
	JwtSecretKey = &cfg.JWTSecret
	return cfg
}
func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
