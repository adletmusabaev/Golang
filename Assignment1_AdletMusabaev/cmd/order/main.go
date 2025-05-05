package main

import (
	"Assignment1_AdletMusabaev/internal/order/handlers"
	"Assignment1_AdletMusabaev/internal/order/repository"
	"Assignment1_AdletMusabaev/internal/order/services"
	"Assignment1_AdletMusabaev/internal/pkg/config"
	"Assignment1_AdletMusabaev/internal/pkg/database"
	"Assignment1_AdletMusabaev/internal/pkg/logger"
	"Assignment1_AdletMusabaev/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logger := logger.NewLogger()

	client, err := database.NewMongoDBConnection(cfg.MongoURI, cfg.MongoDB)
	if err != nil {
		logger.Fatal("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(nil)

	db := client.Database(cfg.MongoDB)
	repo := repository.NewMongoRepository(db, "orders")
	service := services.NewOrderService(repo)
	grpcServer := handlers.NewOrderGRPCServer(service)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		logger.Fatal("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	proto.RegisterOrderServiceServer(server, grpcServer)

	logger.Info("Order Service (gRPC) running on :50052")
	if err := server.Serve(lis); err != nil {
		logger.Fatal("Failed to serve: %v", err)
	}
}
