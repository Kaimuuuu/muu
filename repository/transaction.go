package repository

import (
	"context"
	"kaimuu/model"
	"time"

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
