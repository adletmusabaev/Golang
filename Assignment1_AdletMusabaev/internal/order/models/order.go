package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrderItem struct {
	ProductID string `bson:"product_id"`
	Quantity  int    `bson:"quantity"`
}

type Order struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID string             `bson:"user_id"`
	Items  []OrderItem        `bson:"items"`
	Status string             `bson:"status"`
	Total  float64            `bson:"total"`
}
