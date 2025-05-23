package config

import (
	"os"
)

type Config struct {
	MongoURI       string
	MongoDB        string
	NATSURL        string
	InventoryPort  string
	OrderPort      string
	StatisticsPort string
	InventoryAddr  string
	OrderAddr      string
	StatisticsAddr string
}

func LoadConfig() (*Config, error) {
	return &Config{
		MongoURI:       getEnv("MONGO_URI", "mongodb+srv://ecommerceuser:123@cluster1.xxkrsq1.mongodb.net/?retryWrites=true&w=majority&appName=Cluster1"),
		MongoDB:        getEnv("MONGO_DB", "ecommerce"),
		NATSURL:        getEnv("NATS_URL", "nats://localhost:4222"),
		InventoryPort:  getEnv("INVENTORY_PORT", ":50051"),
		OrderPort:      getEnv("ORDER_PORT", ":50052"),
		StatisticsPort: getEnv("STATISTICS_PORT", ":50053"),
		InventoryAddr:  getEnv("INVENTORY_ADDR", "localhost:50051"),
		OrderAddr:      getEnv("ORDER_ADDR", "localhost:50052"),
		StatisticsAddr: getEnv("STATISTICS_ADDR", "localhost:50053"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
