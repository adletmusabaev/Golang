package grpc

import (
	"Assignment1_AdletMusabaev/internal/statistics/services"
	"Assignment1_AdletMusabaev/proto"
	"context"
)

type StatisticsGRPCServer struct {
	svc *services.StatisticsService
	proto.UnimplementedStatisticsServiceServer
}

func NewStatisticsGRPCServer(svc *services.StatisticsService) *StatisticsGRPCServer {
	return &StatisticsGRPCServer{svc: svc}
}

func (s *StatisticsGRPCServer) GetUserOrdersStatistics(ctx context.Context, req *proto.UserOrderStatisticsRequest) (*proto.UserOrderStatisticsResponse, error) {
	return s.svc.GetUserOrdersStatistics(ctx, req)
}

func (s *StatisticsGRPCServer) GetUserStatistics(ctx context.Context, req *proto.UserStatisticsRequest) (*proto.UserStatisticsResponse, error) {
	return s.svc.GetUserStatistics(ctx, req)
}
