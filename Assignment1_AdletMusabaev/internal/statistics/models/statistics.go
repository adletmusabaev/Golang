package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrderStatistic struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	UserID        string             `bson:"user_id"`
	OrderCount    int32              `bson:"order_count"`
	PeakOrderTime string             `bson:"peak_order_time"`
}

type UserStatistic struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      string             `bson:"user_id"`
	TotalOrders int32              `bson:"total_orders"`
}
