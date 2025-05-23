package main

import (
	"Assignment1_AdletMusabaev/internal/pkg/config"
	"Assignment1_AdletMusabaev/internal/pkg/database"
	"Assignment1_AdletMusabaev/internal/pkg/logger"
	grpcHandlers "Assignment1_AdletMusabaev/internal/statistics/handlers/grpc"
	"Assignment1_AdletMusabaev/internal/statistics/repository"
	"Assignment1_AdletMusabaev/internal/statistics/services"
	"Assignment1_AdletMusabaev/proto"
	"log"
	"net"

	natsLib "github.com/nats-io/nats.go"
	grpcLib "google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logger := logger.NewLogger()

	nc, err := natsLib.Connect(cfg.NATSURL)
	if err != nil {
		logger.Fatal("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	client, err := database.NewMongoDBConnection(cfg.MongoURI, cfg.MongoDB)
	if err != nil {
		logger.Fatal("Failed to initialize MongoDB connection: %v", err)
	}
	defer client.Disconnect(nil)
	db := client.Database(cfg.MongoDB)

	repo := repository.NewStatisticsRepository(db)
	svc := services.NewStatisticsService(repo, nc)
	grpcServer := grpcHandlers.NewStatisticsGRPCServer(svc)

	lis, err := net.Listen("tcp", cfg.StatisticsPort)
	if err != nil {
		logger.Fatal("Failed to listen: %v", err)
	}
	s := grpcLib.NewServer()
	proto.RegisterStatisticsServiceServer(s, grpcServer)

	logger.Info("Starting Statistics microservice gRPC server on %s", cfg.StatisticsPort)
	if err := s.Serve(lis); err != nil {
		logger.Fatal("Failed to serve: %v", err)
	}
}
