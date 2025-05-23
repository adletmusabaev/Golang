package handlers

import (
	"Assignment1_AdletMusabaev/proto"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	clients *GRPCClients
}

func NewHandler(clients *GRPCClients) *Handler {
	return &Handler{clients: clients}
}

func (h *Handler) CreateProduct(c *gin.Context) {
	var req proto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.clients.InventoryClient.CreateProduct(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.clients.InventoryClient.GetProductByID(context.Background(), &proto.GetProductRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var req proto.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Id = id
	resp, err := h.clients.InventoryClient.UpdateProduct(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	_, err := h.clients.InventoryClient.DeleteProduct(context.Background(), &proto.DeleteProductRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func (h *Handler) GetAllProducts(c *gin.Context) {
	resp, err := h.clients.InventoryClient.ListProducts(context.Background(), &proto.ListProductsRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var req proto.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.clients.OrderClient.CreateOrder(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.clients.OrderClient.GetOrderByID(context.Background(), &proto.GetOrderRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.clients.OrderClient.UpdateOrderStatus(context.Background(), &proto.UpdateOrderStatusRequest{
		Id:     id,
		Status: req.Status,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetOrdersByUser(c *gin.Context) {
	userID := c.Param("user_id")
	resp, err := h.clients.OrderClient.ListUserOrders(context.Background(), &proto.ListOrdersRequest{UserId: userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUserOrdersStatistics(c *gin.Context) {
	userID := c.Param("user_id")
	resp, err := h.clients.StatisticsClient.GetUserOrdersStatistics(context.Background(), &proto.UserOrderStatisticsRequest{UserId: userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
