package services

import (
	"Assignment1_AdletMusabaev/internal/inventory/models"
	"Assignment1_AdletMusabaev/internal/pkg/logger"
	"Assignment1_AdletMusabaev/proto"
	"encoding/json"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/patrickmn/go-cache"
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
	natsConn *nats.Conn
	cache    *cache.Cache
	mu       sync.Mutex
	logger   *logger.Logger
}

func NewInventoryService(repo interface {
	Create(product *models.Product) error
	GetByID(id primitive.ObjectID) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id primitive.ObjectID) error
	GetAll() ([]*models.Product, error)
}, natsConn *nats.Conn, logger *logger.Logger) *InventoryService {
	c := cache.New(12*time.Hour, 12*time.Hour)

	svc := &InventoryService{
		repo:     repo,
		natsConn: natsConn,
		cache:    c,
		logger:   logger,
	}

	svc.initializeCache()
	go svc.startCacheRefresh()

	return svc
}

func (s *InventoryService) initializeCache() {
	s.mu.Lock()
	defer s.mu.Unlock()

	products, err := s.repo.GetAll()
	if err != nil {
		s.logger.Error("Failed to initialize cache", err) // Используем Error с правильной сигнатурой
		return
	}

	for _, product := range products {
		s.cache.Set(product.ID.Hex(), product, cache.DefaultExpiration)
	}
	s.logger.Info("Cache initialized with %d items", len(products))
}

func (s *InventoryService) startCacheRefresh() {
	ticker := time.NewTicker(12 * time.Hour)
	for {
		select {
		case <-ticker.C:
			s.logger.Info("Refreshing cache")
			s.initializeCache()
		}
	}
}

func (s *InventoryService) CreateProduct(product *models.Product) error {
	err := s.repo.Create(product)
	if err != nil {
		s.logger.Error("Failed to create product", err)
		return err
	}

	s.mu.Lock()
	s.cache.Set(product.ID.Hex(), product, cache.DefaultExpiration)
	s.mu.Unlock()

	event := &proto.InventoryEvent{
		Id:        product.ID.Hex(),
		Action:    "created",
		Timestamp: time.Now().Unix(),
	}
	data, err := json.Marshal(event)
	if err != nil {
		s.logger.Error("Failed to marshal inventory event", err)
		return err
	}
	if err := s.natsConn.Publish("inventory.created", data); err != nil {
		s.logger.Error("Failed to publish inventory.created event", err)
		return err
	}
	return nil
}

func (s *InventoryService) UpdateProduct(product *models.Product) error {
	err := s.repo.Update(product)
	if err != nil {
		s.logger.Error("Failed to update product", err)
		return err
	}

	s.mu.Lock()
	s.cache.Set(product.ID.Hex(), product, cache.DefaultExpiration)
	s.mu.Unlock()

	event := &proto.InventoryEvent{
		Id:        product.ID.Hex(),
		Action:    "updated",
		Timestamp: time.Now().Unix(),
	}
	data, err := json.Marshal(event)
	if err != nil {
		s.logger.Error("Failed to marshal inventory event", err)
		return err
	}
	if err := s.natsConn.Publish("inventory.updated", data); err != nil {
		s.logger.Error("Failed to publish inventory.updated event", err)
		return err
	}
	return nil
}

func (s *InventoryService) DeleteProduct(id primitive.ObjectID) error {
	err := s.repo.Delete(id)
	if err != nil {
		s.logger.Error("Failed to delete product", err)
		return err
	}

	s.mu.Lock()
	s.cache.Delete(id.Hex())
	s.mu.Unlock()

	event := &proto.InventoryEvent{
		Id:        id.Hex(),
		Action:    "deleted",
		Timestamp: time.Now().Unix(),
	}
	data, err := json.Marshal(event)
	if err != nil {
		s.logger.Error("Failed to marshal inventory event", err)
		return err
	}
	if err := s.natsConn.Publish("inventory.deleted", data); err != nil {
		s.logger.Error("Failed to publish inventory.deleted event", err)
		return err
	}
	return nil
}

func (s *InventoryService) GetProductByID(id primitive.ObjectID) (*models.Product, error) {
	s.mu.Lock()
	if cachedItem, found := s.cache.Get(id.Hex()); found {
		s.mu.Unlock()
		return cachedItem.(*models.Product), nil
	}
	s.mu.Unlock()

	product, err := s.repo.GetByID(id)
	if err != nil {
		s.logger.Error("Failed to get product by ID", err)
		return nil, err
	}

	s.mu.Lock()
	s.cache.Set(product.ID.Hex(), product, cache.DefaultExpiration)
	s.mu.Unlock()

	return product, nil
}

func (s *InventoryService) GetAllProducts() ([]*models.Product, error) {
	s.mu.Lock()
	items := s.cache.Items()
	if len(items) > 0 {
		products := make([]*models.Product, 0, len(items))
		for _, item := range items {
			products = append(products, item.Object.(*models.Product))
		}
		s.mu.Unlock()
		return products, nil
	}
	s.mu.Unlock()

	products, err := s.repo.GetAll()
	if err != nil {
		s.logger.Error("Failed to get all products", err)
		return nil, err
	}

	s.mu.Lock()
	for _, product := range products {
		s.cache.Set(product.ID.Hex(), product, cache.DefaultExpiration)
	}
	s.mu.Unlock()

	return products, nil
}
