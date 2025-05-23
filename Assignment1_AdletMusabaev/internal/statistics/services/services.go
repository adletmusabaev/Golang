package services

import (
	"Assignment1_AdletMusabaev/internal/statistics/models"
	"Assignment1_AdletMusabaev/internal/statistics/repository"
	"Assignment1_AdletMusabaev/proto"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type StatisticsService struct {
	repo     *repository.StatisticsRepository
	natsConn *nats.Conn
}

func NewStatisticsService(repo *repository.StatisticsRepository, natsConn *nats.Conn) *StatisticsService {
	svc := &StatisticsService{
		repo:     repo,
		natsConn: natsConn,
	}
	if natsConn != nil {
		svc.subscribeToNATS()
	}
	return svc
}

func (s *StatisticsService) subscribeToNATS() {
	s.natsConn.Subscribe("order.created", func(msg *nats.Msg) {
		var event proto.OrderEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Printf("Failed to unmarshal order event: %v", err)
			return
		}
		if err := s.ProcessOrderEvent(&event); err != nil {
			log.Printf("Failed to process order event: %v", err)
		}
	})

	s.natsConn.Subscribe("inventory.created", func(msg *nats.Msg) {
		var event proto.InventoryEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Printf("Failed to unmarshal inventory event: %v", err)
			return
		}
		// Process inventory event if needed
	})
}

func (s *StatisticsService) GetUserOrdersStatistics(ctx context.Context, req *proto.UserOrderStatisticsRequest) (*proto.UserOrderStatisticsResponse, error) {
	stat, err := s.repo.GetOrderStatistic(ctx, req.UserId)
	if err != nil {
		return &proto.UserOrderStatisticsResponse{
			OrderCount:    0,
			PeakOrderTime: "",
		}, nil
	}
	return &proto.UserOrderStatisticsResponse{
		OrderCount:    stat.OrderCount,
		PeakOrderTime: stat.PeakOrderTime,
	}, nil
}

func (s *StatisticsService) GetUserStatistics(ctx context.Context, req *proto.UserStatisticsRequest) (*proto.UserStatisticsResponse, error) {
	stat, err := s.repo.GetUserStatistic(ctx, req.UserId)
	if err != nil {
		return &proto.UserStatisticsResponse{
			TotalUsers:     1,
			UserOrderCount: 0,
		}, nil
	}
	return &proto.UserStatisticsResponse{
		TotalUsers:     1,
		UserOrderCount: stat.TotalOrders,
	}, nil
}

func (s *StatisticsService) ProcessOrderEvent(event *proto.OrderEvent) error {
	stat, err := s.repo.GetOrderStatistic(context.Background(), event.UserId)
	if err != nil {
		stat = &models.OrderStatistic{UserID: event.UserId}
	}

	if event.Action == "created" {
		stat.OrderCount++
		t := time.Unix(event.Timestamp, 0)
		stat.PeakOrderTime = t.Format("15:00") + "-" + t.Add(time.Hour).Format("15:00")
	}
	return s.repo.SaveOrderStatistic(context.Background(), stat)
}
