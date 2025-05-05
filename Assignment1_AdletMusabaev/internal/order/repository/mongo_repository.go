package repository

import (
	"Assignment1_AdletMusabaev/internal/order/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database, collectionName string) *MongoRepository {
	return &MongoRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *MongoRepository) Create(order *models.Order) error {
	_, err := r.collection.InsertOne(context.Background(), order)
	return err
}

func (r *MongoRepository) GetByID(id primitive.ObjectID) (*models.Order, error) {
	var order models.Order
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&order)
	return &order, err
}

func (r *MongoRepository) UpdateStatus(id primitive.ObjectID, status string) error {
	_, err := r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"status": status}},
	)
	return err
}

func (r *MongoRepository) GetByUserID(userID string) ([]*models.Order, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var orders []*models.Order
	for cursor.Next(context.Background()) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}
