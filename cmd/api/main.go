package main

import (
	"log"
	"social-backend/internal/config"
	"social-backend/internal/handler"
	"social-backend/internal/repository"
	"social-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect DB
	config.ConnectDB()

	// dependency injection
	repo := repository.NewUserRepository(config.DB)
	authService := service.NewAuthService(repo)
	authHandler := handler.NewAuthHandler(authService)

	r := gin.Default()

	r.POST("/register", authHandler.Register)

	if err := r.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
