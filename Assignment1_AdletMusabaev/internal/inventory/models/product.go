package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Category string             `bson:"category"`
	Stock    int                `bson:"stock"`
	Price    float64            `bson:"price"`
}
