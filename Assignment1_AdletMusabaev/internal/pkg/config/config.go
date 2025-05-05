package config

import (
	"os"
)

type Config struct {
	MongoURI string
	MongoDB  string
}

func LoadConfig() (*Config, error) {
	return &Config{
		MongoURI: getEnv("MONGO_URI", "mongodb+srv://ecommerceuser:123@cluster1.xxkrsq1.mongodb.net/?retryWrites=true&w=majority&appName=Cluster1"),
		MongoDB:  getEnv("MONGO_DB", "ecommerce"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
