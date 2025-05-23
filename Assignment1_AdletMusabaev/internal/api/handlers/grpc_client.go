package handlers

import (
	"Assignment1_AdletMusabaev/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClients struct {
	InventoryClient  proto.InventoryServiceClient
	OrderClient      proto.OrderServiceClient
	StatisticsClient proto.StatisticsServiceClient
	inventoryConn    *grpc.ClientConn // Храним соединение для закрытия
	orderConn        *grpc.ClientConn // Храним соединение для закрытия
	statisticsConn   *grpc.ClientConn // Храним соединение для закрытия
}

func NewGRPCClients(inventoryAddr, orderAddr, statisticsAddr string) (*GRPCClients, error) {
	// Подключение к Inventory Service
	inventoryConn, err := grpc.Dial(inventoryAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// Подключение к Order Service
	orderConn, err := grpc.Dial(orderAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		inventoryConn.Close()
		return nil, err
	}

	// Подключение к Statistics Service
	statisticsConn, err := grpc.Dial(statisticsAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		inventoryConn.Close()
		orderConn.Close()
		return nil, err
	}

	return &GRPCClients{
		InventoryClient:  proto.NewInventoryServiceClient(inventoryConn),
		OrderClient:      proto.NewOrderServiceClient(orderConn),
		StatisticsClient: proto.NewStatisticsServiceClient(statisticsConn),
		inventoryConn:    inventoryConn,
		orderConn:        orderConn,
		statisticsConn:   statisticsConn,
	}, nil
}

func (c *GRPCClients) Close() {
	if c.inventoryConn != nil {
		c.inventoryConn.Close()
	}
	if c.orderConn != nil {
		c.orderConn.Close()
	}
	if c.statisticsConn != nil {
		c.statisticsConn.Close()
	}
}
