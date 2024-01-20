package repository

import (
	"context"
	"kaimuu/model"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderRepository struct {
	col mongo.Collection
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	return &OrderRepository{
		col: *db.Collection("orders"),
	}
}

func (or *OrderRepository) Insert(o *model.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := or.col.InsertOne(ctx, o)
	if err != nil {
		return err
	}

	return nil
}

func (or *OrderRepository) Update(orderId string, o *model.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := or.col.UpdateOne(ctx, bson.D{{"id", orderId}}, bson.D{{"$set", o}})
	if result.MatchedCount == 0 {
		return errors.Errorf("invalid order id '%s'", orderId)
	}
	if err != nil {
		return err
	}

	return nil
}

func (or *OrderRepository) Delete(orderId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := or.col.DeleteOne(ctx, bson.D{{"id", orderId}})
	if result.DeletedCount == 0 {
		return errors.Errorf("invalid order id '%s'", orderId)
	}
	if err != nil {
		return err
	}

	return nil
}

func (or *OrderRepository) GetById(orderId string) (*model.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var o model.Order
	err := or.col.FindOne(ctx, bson.D{{"id", orderId}}).Decode(&o)
	if err != nil {
		return &model.Order{}, err
	}

	return &o, nil
}

func (or *OrderRepository) GetPendingOrder() ([]model.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := or.col.Find(ctx, bson.D{{"status", model.Pending}})
	if err != nil {
		return []model.Order{}, err
	}

	var orders []model.Order
	if err := cur.All(context.Background(), &orders); err != nil {
		return []model.Order{}, err
	}

	return orders, nil
}

func (or *OrderRepository) GetOrderByToken(token string) ([]model.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Find().SetSort(bson.D{{"createdAt", -1}})

	cur, err := or.col.Find(ctx, bson.D{{"orderBy", token}}, opts)
	if err != nil {
		return []model.Order{}, err
	}

	var orders []model.Order
	if err := cur.All(context.Background(), &orders); err != nil {
		return []model.Order{}, err
	}

	return orders, nil
}
