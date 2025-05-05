package handlers

import (
	"Assignment1_AdletMusabaev/internal/inventory/models"
	"Assignment1_AdletMusabaev/internal/inventory/services"
	"Assignment1_AdletMusabaev/proto"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryGRPCServer struct {
	proto.UnimplementedInventoryServiceServer
	service *services.InventoryService
}

func NewInventoryGRPCServer(service *services.InventoryService) *InventoryGRPCServer {
	return &InventoryGRPCServer{service: service}
}

func (s *InventoryGRPCServer) CreateProduct(ctx context.Context, req *proto.CreateProductRequest) (*proto.ProductResponse, error) {
	product := &models.Product{
		Name:     req.Name,
		Category: req.Category,
		Stock:    int(req.Stock),
		Price:    req.Price,
	}
	if err := s.service.CreateProduct(product); err != nil {
		return nil, err
	}
	return &proto.ProductResponse{
		Id:       product.ID.Hex(),
		Name:     product.Name,
		Category: product.Category,
		Stock:    int32(product.Stock),
		Price:    product.Price,
	}, nil
}

func (s *InventoryGRPCServer) GetProductByID(ctx context.Context, req *proto.GetProductRequest) (*proto.ProductResponse, error) {
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}
	product, err := s.service.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return &proto.ProductResponse{
		Id:       product.ID.Hex(),
		Name:     product.Name,
		Category: product.Category,
		Stock:    int32(product.Stock),
		Price:    product.Price,
	}, nil
}

func (s *InventoryGRPCServer) UpdateProduct(ctx context.Context, req *proto.UpdateProductRequest) (*proto.ProductResponse, error) {
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}
	product := &models.Product{
		ID:       id,
		Name:     req.Name,
		Category: req.Category,
		Stock:    int(req.Stock),
		Price:    req.Price,
	}
	if err := s.service.UpdateProduct(product); err != nil {
		return nil, err
	}
	return &proto.ProductResponse{
		Id:       product.ID.Hex(),
		Name:     product.Name,
		Category: product.Category,
		Stock:    int32(product.Stock),
		Price:    product.Price,
	}, nil
}

func (s *InventoryGRPCServer) DeleteProduct(ctx context.Context, req *proto.DeleteProductRequest) (*proto.Empty, error) {
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}
	if err := s.service.DeleteProduct(id); err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}

func (s *InventoryGRPCServer) ListProducts(ctx context.Context, req *proto.ListProductsRequest) (*proto.ListProductsResponse, error) {
	products, err := s.service.GetAllProducts()
	if err != nil {
		return nil, err
	}
	protoProducts := make([]*proto.ProductResponse, len(products))
	for i, p := range products {
		protoProducts[i] = &proto.ProductResponse{
			Id:       p.ID.Hex(),
			Name:     p.Name,
			Category: p.Category,
			Stock:    int32(p.Stock),
			Price:    p.Price,
		}
	}
	return &proto.ListProductsResponse{Products: protoProducts}, nil
}
