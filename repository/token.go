package repository

import (
	"context"
	"kaimuu/model"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenRepository struct {
	col *mongo.Collection
}

func NewTokenRepository(db *mongo.Database) *TokenRepository {
	return &TokenRepository{
		col: db.Collection("tokens"),
	}
}

func (tr *TokenRepository) Get(token string) (*model.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var c model.Client
	err := tr.col.FindOne(ctx, bson.D{{"token", token}}).Decode(&c)
	if err != nil {
		return &model.Client{}, err
	}

	return &c, nil
}

func (tr *TokenRepository) Delete(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := tr.col.DeleteOne(ctx, bson.D{{"token", token}})
	if result.DeletedCount == 0 {
		return errors.Errorf("invalid token '%s'", token)
	}
	if err != nil {
		return err
	}

	return nil
}

func (tr *TokenRepository) All() ([]model.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := tr.col.Find(ctx, bson.D{})
	if err != nil {
		return []model.Client{}, err
	}

	var clients []model.Client
	if err := cur.All(context.Background(), &clients); err != nil {
		return []model.Client{}, err
	}

	return clients, nil
}

func (tr *TokenRepository) Insert(c *model.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := tr.col.InsertOne(ctx, c)
	if err != nil {
		return err
	}

	return nil
}
