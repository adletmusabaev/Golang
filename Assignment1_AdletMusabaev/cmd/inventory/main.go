package main

import (
	"Assignment1_AdletMusabaev/internal/inventory/handlers"
	"Assignment1_AdletMusabaev/internal/inventory/repository"
	"Assignment1_AdletMusabaev/internal/inventory/services"
	"Assignment1_AdletMusabaev/internal/pkg/config"
	"Assignment1_AdletMusabaev/internal/pkg/database"
	"Assignment1_AdletMusabaev/internal/pkg/logger"
	"Assignment1_AdletMusabaev/proto"
	"log"
	"net"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logger := logger.NewLogger()

	// Подключение к MongoDB
	client, err := database.NewMongoDBConnection(cfg.MongoURI, cfg.MongoDB)
	if err != nil {
		logger.Fatal("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(nil)

	// Подключение к NATS
	nc, err := nats.Connect(cfg.NATSURL)
	if err != nil {
		logger.Fatal("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	db := client.Database(cfg.MongoDB)
	repo := repository.NewMongoRepository(db, "products")
	service := services.NewInventoryService(repo, nc, logger) // Передаем logger
	grpcServer := handlers.NewInventoryGRPCServer(service)

	lis, err := net.Listen("tcp", cfg.InventoryPort)
	if err != nil {
		logger.Fatal("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	proto.RegisterInventoryServiceServer(server, grpcServer)

	logger.Info("Inventory Service (gRPC) running on %s", cfg.InventoryPort)
	if err := server.Serve(lis); err != nil {
		logger.Fatal("Failed to serve: %v", err)
	}
}
