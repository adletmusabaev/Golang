package handlers

import (
	"Assignment1_AdletMusabaev/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClients struct {
	InventoryClient proto.InventoryServiceClient
	OrderClient     proto.OrderServiceClient
}

func NewGRPCClients(inventoryAddr, orderAddr string) (*GRPCClients, error) {
	inventoryConn, err := grpc.Dial(inventoryAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	orderConn, err := grpc.Dial(orderAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		inventoryConn.Close()
		return nil, err
	}
	return &GRPCClients{
		InventoryClient: proto.NewInventoryServiceClient(inventoryConn),
		OrderClient:     proto.NewOrderServiceClient(orderConn),
	}, nil
}
