package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"crud/internal/config"
	service "crud/internal/crud"
	"crud/internal/handler"
	"crud/internal/repository"
	"crud/internal/routes"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	db := client.Database(cfg.MongoDBName)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandlers := handler.NewUserHandlers(userService)

	r := gin.Default()
	routes.RegisterRoutes(r, userHandlers)

	port := cfg.AppPort
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server running on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
