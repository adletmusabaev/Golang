package services

import (
	"Assignment1_AdletMusabaev/internal/order/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService struct {
	repo interface {
		Create(order *models.Order) error
		GetByID(id primitive.ObjectID) (*models.Order, error)
		UpdateStatus(id primitive.ObjectID, status string) error
		GetByUserID(userID string) ([]*models.Order, error)
	}
}

func NewOrderService(repo interface {
	Create(order *models.Order) error
	GetByID(id primitive.ObjectID) (*models.Order, error)
	UpdateStatus(id primitive.ObjectID, status string) error
	GetByUserID(userID string) ([]*models.Order, error)
}) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	return s.repo.Create(order)
}

func (s *OrderService) GetOrderByID(id primitive.ObjectID) (*models.Order, error) {
	return s.repo.GetByID(id)
}

func (s *OrderService) UpdateOrderStatus(id primitive.ObjectID, status string) error {
	return s.repo.UpdateStatus(id, status)
}

func (s *OrderService) GetOrdersByUserID(userID string) ([]*models.Order, error) {
	return s.repo.GetByUserID(userID)
}
