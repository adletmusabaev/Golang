package services

import (
	"Assignment1_AdletMusabaev/internal/order/models"
	"Assignment1_AdletMusabaev/proto"
	"encoding/json"
	"time"

	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService struct {
	repo interface {
		Create(order *models.Order) error
		GetByID(id primitive.ObjectID) (*models.Order, error)
		UpdateStatus(id primitive.ObjectID, status string) error
		GetByUserID(userID string) ([]*models.Order, error)
	}
	natsConn *nats.Conn
}

func NewOrderService(repo interface {
	Create(order *models.Order) error
	GetByID(id primitive.ObjectID) (*models.Order, error)
	UpdateStatus(id primitive.ObjectID, status string) error
	GetByUserID(userID string) ([]*models.Order, error)
}, natsConn *nats.Conn) *OrderService {
	return &OrderService{repo: repo, natsConn: natsConn}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	err := s.repo.Create(order)
	if err != nil {
		return err
	}
	// Публикация события в NATS
	event := &proto.OrderEvent{
		Id:        order.ID.Hex(),
		UserId:    order.UserID,
		Action:    "created",
		Timestamp: time.Now().Unix(),
	}
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return s.natsConn.Publish("order.created", data)
}

func (s *OrderService) UpdateOrderStatus(id primitive.ObjectID, status string) error {
	err := s.repo.UpdateStatus(id, status)
	if err != nil {
		return err
	}
	// Публикация события в NATS
	event := &proto.OrderEvent{
		Id:        id.Hex(),
		Action:    "updated",
		Timestamp: time.Now().Unix(),
	}
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return s.natsConn.Publish("order.updated", data)
}

func (s *OrderService) GetOrderByID(id primitive.ObjectID) (*models.Order, error) {
	return s.repo.GetByID(id)
}

func (s *OrderService) GetOrdersByUserID(userID string) ([]*models.Order, error) {
	return s.repo.GetByUserID(userID)
}
