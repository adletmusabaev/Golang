package repository

import (
	"Assignment1_AdletMusabaev/internal/inventory/models"
	"context"
	"log"
	"time"

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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Убедимся, что ID не задан, чтобы MongoDB сгенерировал его
	if product.ID.IsZero() {
		product.ID = primitive.NewObjectID()
	}

	result, err := r.collection.InsertOne(ctx, product)
	if err != nil {
		log.Printf("Error inserting product: %v", err)
		return err
	}

	// Обновляем ID в объекте product
	if insertedID, ok := result.InsertedID.(primitive.ObjectID); ok {
		product.ID = insertedID
	}

	log.Printf("Product inserted successfully: %v", product.ID.Hex())
	return nil
}

func (r *MongoRepository) GetByID(id primitive.ObjectID) (*models.Product, error) {
	var product models.Product
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	return &product, err
}

func (r *MongoRepository) Update(product *models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.UpdateOne(
		ctx,
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *MongoRepository) GetAll() ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*models.Product
	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}
