package handlers

import (
	"Assignment1_AdletMusabaev/internal/order/models"
	"Assignment1_AdletMusabaev/internal/order/services"
	"Assignment1_AdletMusabaev/proto"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderGRPCServer struct {
	proto.UnimplementedOrderServiceServer
	service *services.OrderService
}

func NewOrderGRPCServer(service *services.OrderService) *OrderGRPCServer {
	return &OrderGRPCServer{service: service}
}

func (s *OrderGRPCServer) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.OrderResponse, error) {
	order := &models.Order{
		UserID: req.UserId,
		Status: "pending",
	}
	for _, item := range req.Items {
		order.Items = append(order.Items, models.OrderItem{
			ProductID: item.ProductId,
			Quantity:  int(item.Quantity),
		})
	}
	order.Total = float64(len(order.Items) * 100) // Simplified total
	if err := s.service.CreateOrder(order); err != nil {
		return nil, err
	}
	protoItems := make([]*proto.OrderItem, len(order.Items))
	for i, item := range order.Items {
		protoItems[i] = &proto.OrderItem{
			ProductId: item.ProductID,
			Quantity:  int32(item.Quantity),
		}
	}
	return &proto.OrderResponse{
		Id:     order.ID.Hex(),
		UserId: order.UserID,
		Items:  protoItems,
		Status: order.Status,
		Total:  order.Total,
	}, nil
}

func (s *OrderGRPCServer) GetOrderByID(ctx context.Context, req *proto.GetOrderRequest) (*proto.OrderResponse, error) {
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}
	order, err := s.service.GetOrderByID(id)
	if err != nil {
		return nil, err
	}
	protoItems := make([]*proto.OrderItem, len(order.Items))
	for i, item := range order.Items {
		protoItems[i] = &proto.OrderItem{
			ProductId: item.ProductID,
			Quantity:  int32(item.Quantity),
		}
	}
	return &proto.OrderResponse{
		Id:     order.ID.Hex(),
		UserId: order.UserID,
		Items:  protoItems,
		Status: order.Status,
		Total:  order.Total,
	}, nil
}

func (s *OrderGRPCServer) UpdateOrderStatus(ctx context.Context, req *proto.UpdateOrderStatusRequest) (*proto.OrderResponse, error) {
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}
	if err := s.service.UpdateOrderStatus(id, req.Status); err != nil {
		return nil, err
	}
	order, err := s.service.GetOrderByID(id)
	if err != nil {
		return nil, err
	}
	protoItems := make([]*proto.OrderItem, len(order.Items))
	for i, item := range order.Items {
		protoItems[i] = &proto.OrderItem{
			ProductId: item.ProductID,
			Quantity:  int32(item.Quantity),
		}
	}
	return &proto.OrderResponse{
		Id:     order.ID.Hex(),
		UserId: order.UserID,
		Items:  protoItems,
		Status: order.Status,
		Total:  order.Total,
	}, nil
}

func (s *OrderGRPCServer) ListUserOrders(ctx context.Context, req *proto.ListOrdersRequest) (*proto.ListOrdersResponse, error) {
	orders, err := s.service.GetOrdersByUserID(req.UserId)
	if err != nil {
		return nil, err
	}
	protoOrders := make([]*proto.OrderResponse, len(orders))
	for i, order := range orders {
		protoItems := make([]*proto.OrderItem, len(order.Items))
		for j, item := range order.Items {
			protoItems[j] = &proto.OrderItem{
				ProductId: item.ProductID,
				Quantity:  int32(item.Quantity),
			}
		}
		protoOrders[i] = &proto.OrderResponse{
			Id:     order.ID.Hex(),
			UserId: order.UserID,
			Items:  protoItems,
			Status: order.Status,
			Total:  order.Total,
		}
	}
	return &proto.ListOrdersResponse{Orders: protoOrders}, nil
}
