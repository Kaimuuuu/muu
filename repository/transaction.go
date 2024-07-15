package repository

import (
	"context"
	"github.com/Kaimuuuu/muu/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepository struct {
	col *mongo.Collection
}

func NewTransactionRepository(db *mongo.Database) *TransactionRepository {
	return &TransactionRepository{
		col: db.Collection("transactions"),
	}
}

func (tr *TransactionRepository) Insert(t *model.Transaction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := tr.col.InsertOne(ctx, t)
	if err != nil {
		return err
	}

	return nil
}

func (tr *TransactionRepository) All() ([]model.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := tr.col.Find(ctx, bson.D{})
	if err != nil {
		return []model.Transaction{}, err
	}

	var transactions []model.Transaction
	if err := cur.All(context.Background(), &transactions); err != nil {
		return []model.Transaction{}, err
	}

	return transactions, nil
}
