package repository

import (
	"context"
	"kaimuu/model"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PromotionRepository struct {
	col *mongo.Collection
}

func NewPromotionRepository(db *mongo.Database) *PromotionRepository {
	return &PromotionRepository{
		col: db.Collection("ptions"),
	}
}

func (pr *PromotionRepository) Insert(p *model.Promotion) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := pr.col.InsertOne(ctx, p)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PromotionRepository) Update(id string, p *model.Promotion) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := pr.col.UpdateOne(ctx, bson.D{{"id", id}}, bson.D{{"$set", p}})
	if result.MatchedCount == 0 {
		return errors.Errorf("invalid ption id '%s'", id)
	}
	if err != nil {
		return err
	}

	return nil
}

func (pr *PromotionRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := pr.col.DeleteOne(ctx, bson.D{{"id", id}})
	if result.DeletedCount == 0 {
		return errors.Errorf("invalid ption id '%s'", id)
	}
	if err != nil {
		return err
	}

	return nil
}

func (pr *PromotionRepository) GetById(id string) (*model.Promotion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var p model.Promotion
	if err := pr.col.FindOne(ctx, bson.D{{"id", id}}).Decode(&p); err != nil {
		return &model.Promotion{}, err
	}

	return &p, nil
}

func (pr *PromotionRepository) GetAll() ([]model.Promotion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := pr.col.Find(ctx, bson.D{})
	if err != nil {
		return []model.Promotion{}, err
	}

	var promotions []model.Promotion
	if err := cur.All(context.Background(), &promotions); err != nil {
		return []model.Promotion{}, err
	}

	return promotions, nil
}
