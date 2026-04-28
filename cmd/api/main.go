package main

import (
	"log"
	"social-backend/internal/config"
	"social-backend/internal/handler"
	"social-backend/internal/middleware"
	"social-backend/internal/repository"
	"social-backend/internal/service"
	"social-backend/pkg/worker"

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

	// init worker pool
	workerPool := worker.NewWorkerPool(5)
	defer workerPool.Stop()

	// dependency injection
	repo := repository.NewUserRepository(config.DB)
	authService := service.NewAuthService(repo, workerPool)
	authHandler := handler.NewAuthHandler(authService)

	r := gin.Default()

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/me", func(c *gin.Context) {
		c.JSON(200, gin.H{"user_id": c.GetString("user_id")})
	})

	if err := r.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
