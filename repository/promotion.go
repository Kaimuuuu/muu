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
		col: db.Collection("promotions"),
	}
}

func (pr *PromotionRepository) Insert(promo *model.Promotion) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := pr.col.InsertOne(ctx, promo)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PromotionRepository) Update(promotionId string, promo *model.Promotion) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := pr.col.UpdateOne(ctx, bson.D{{"id", promotionId}}, bson.D{{"$set", promo}})
	if result.MatchedCount == 0 {
		return errors.Errorf("invalid promotion id '%s'", promotionId)
	}
	if err != nil {
		return err
	}

	return nil
}

func (pr *PromotionRepository) Delete(promotionId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := pr.col.DeleteOne(ctx, bson.D{{"id", promotionId}})
	if result.DeletedCount == 0 {
		return errors.Errorf("invalid promotion id '%s'", promotionId)
	}
	if err != nil {
		return err
	}

	return nil
}

func (pr *PromotionRepository) GetById(promotionId string) (*model.Promotion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var promo model.Promotion
	if err := pr.col.FindOne(ctx, bson.D{{"id", promotionId}}).Decode(&promo); err != nil {
		return &model.Promotion{}, err
	}

	return &promo, nil
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
