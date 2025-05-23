package routes

import (
	"Assignment1_AdletMusabaev/internal/api/handlers"
	"Assignment1_AdletMusabaev/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *handlers.Handler) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())

	r.POST("/products", handler.CreateProduct)
	r.GET("/products/:id", handler.GetProduct)
	r.PUT("/products/:id", handler.UpdateProduct)
	r.DELETE("/products/:id", handler.DeleteProduct)
	r.GET("/products", handler.GetAllProducts)

	r.POST("/orders", handler.CreateOrder)
	r.GET("/orders/:id", handler.GetOrder)
	r.PUT("/orders/:id/status", handler.UpdateOrderStatus)
	r.GET("/users/:user_id/orders", handler.GetOrdersByUser)

	r.GET("/users/:user_id/order-statistics", handler.GetUserOrdersStatistics) // Добавляем маршрут

	return r
}
