package main

import (
	"fmt"
	"log"
	"os"

	"github.com/erfanyousefi/codenight/config"
	"github.com/erfanyousefi/codenight/internal/auth"
	"github.com/erfanyousefi/codenight/internal/category"
	mainPage "github.com/erfanyousefi/codenight/internal/main-page"
	"github.com/erfanyousefi/codenight/internal/user"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	config.InitDB(&cfg)
	if err := config.GORMDB.AutoMigrate(&user.User{}, &category.Category{}); err != nil {
		fmt.Println("Cannot run migration")
		fmt.Print(err)
	}
	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		log.Fatal("cannot create uploads folder")
	}
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Static("/uploads", "./uploads")
	auth.RegisterRoutes(r, cfg.JWTSecret)
	category.RegisterRoutes(r)
	mainPage.RegisterRoutes(r)
	server := fmt.Sprintf(":%s", cfg.ServerPort)
	fmt.Println("http://localhost:8080")
	r.Run(server)
}
