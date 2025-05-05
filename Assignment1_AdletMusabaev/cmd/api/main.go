package main

import (
	"Assignment1_AdletMusabaev/internal/api/handlers"
	"Assignment1_AdletMusabaev/internal/api/routes"
	"Assignment1_AdletMusabaev/internal/pkg/config"
	"Assignment1_AdletMusabaev/internal/pkg/logger"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logger := logger.NewLogger()
	logger.Info("Loaded config: MongoURI=%s, MongoDB=%s", cfg.MongoURI, cfg.MongoDB)

	grpcClients, err := handlers.NewGRPCClients("localhost:50051", "localhost:50052")
	if err != nil {
		logger.Fatal("Failed to initialize gRPC clients: %v", err)
	}

	handler := handlers.NewHandler(grpcClients)
	router := routes.SetupRouter(handler)

	logger.Info("API Gateway running on :8080")
	if err := router.Run(":8080"); err != nil {
		logger.Fatal("Failed to run server: %v", err)
	}
}
