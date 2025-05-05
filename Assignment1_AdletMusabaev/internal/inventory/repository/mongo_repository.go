package repository

import (
	"Assignment1_AdletMusabaev/internal/inventory/models"
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

func (r *MongoRepository) Create(product *models.Product) error {
	_, err := r.collection.InsertOne(context.Background(), product)
	return err
}

func (r *MongoRepository) GetByID(id primitive.ObjectID) (*models.Product, error) {
	var product models.Product
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&product)
	return &product, err
}

func (r *MongoRepository) Update(product *models.Product) error {
	_, err := r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": product.ID},
		bson.M{"$set": bson.M{
			"name":     product.Name,
			"category": product.Category,
			"stock":    product.Stock,
			"price":    product.Price,
		}},
	)
	return err
}

func (r *MongoRepository) Delete(id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (r *MongoRepository) GetAll() ([]*models.Product, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var products []*models.Product
	for cursor.Next(context.Background()) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}
