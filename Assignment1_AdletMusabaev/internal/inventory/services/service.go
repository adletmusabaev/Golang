package services

import (
	"Assignment1_AdletMusabaev/internal/inventory/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryService struct {
	repo interface {
		Create(product *models.Product) error
		GetByID(id primitive.ObjectID) (*models.Product, error)
		Update(product *models.Product) error
		Delete(id primitive.ObjectID) error
		GetAll() ([]*models.Product, error)
	}
}

func NewInventoryService(repo interface {
	Create(product *models.Product) error
	GetByID(id primitive.ObjectID) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id primitive.ObjectID) error
	GetAll() ([]*models.Product, error)
}) *InventoryService {
	return &InventoryService{repo: repo}
}

func (s *InventoryService) CreateProduct(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *InventoryService) GetProductByID(id primitive.ObjectID) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *InventoryService) UpdateProduct(product *models.Product) error {
	return s.repo.Update(product)
}

func (s *InventoryService) DeleteProduct(id primitive.ObjectID) error {
	return s.repo.Delete(id)
}

func (s *InventoryService) GetAllProducts() ([]*models.Product, error) {
	return s.repo.GetAll()
}
