package repository

import (
	"Assignment1_AdletMusabaev/internal/statistics/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StatisticsRepository struct {
	db *mongo.Database
}

func NewStatisticsRepository(db *mongo.Database) *StatisticsRepository {
	return &StatisticsRepository{db: db}
}

func (r *StatisticsRepository) SaveOrderStatistic(ctx context.Context, stat *models.OrderStatistic) error {
	_, err := r.db.Collection("order_statistics").InsertOne(ctx, stat)
	return err
}

func (r *StatisticsRepository) GetOrderStatistic(ctx context.Context, userID string) (*models.OrderStatistic, error) {
	var stat models.OrderStatistic
	err := r.db.Collection("order_statistics").FindOne(ctx, bson.M{"user_id": userID}).Decode(&stat)
	return &stat, err
}

func (r *StatisticsRepository) SaveUserStatistic(ctx context.Context, stat *models.UserStatistic) error {
	_, err := r.db.Collection("user_statistics").InsertOne(ctx, stat)
	return err
}

func (r *StatisticsRepository) GetUserStatistic(ctx context.Context, userID string) (*models.UserStatistic, error) {
	var stat models.UserStatistic
	err := r.db.Collection("user_statistics").FindOne(ctx, bson.M{"user_id": userID}).Decode(&stat)
	return &stat, err
}
