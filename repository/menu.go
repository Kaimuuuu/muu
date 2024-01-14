package repository

import (
	"context"
	"fmt"
	"kaimuu/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MenuRepository struct {
	col *mongo.Collection
}

func NewMenuRepository(db *mongo.Database) *MenuRepository {
	return &MenuRepository{
		col: db.Collection("menus"),
	}
}

func (mr *MenuRepository) Insert(m *model.MenuItem) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := mr.col.InsertOne(ctx, m)
	if err != nil {
		return err
	}

	return nil
}

func (mr *MenuRepository) Update(menuItemId string, m *model.MenuItem) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := mr.col.UpdateOne(ctx, bson.D{{"id", menuItemId}}, bson.D{{"$set", m}})
	if result.MatchedCount == 0 {
		return fmt.Errorf("invalid menu id {%s}", menuItemId)
	}
	if err != nil {
		return err
	}

	return nil
}

func (mr *MenuRepository) Delete(menuItemId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := mr.col.DeleteOne(ctx, bson.D{{"id", menuItemId}})
	if result.DeletedCount == 0 {
		return fmt.Errorf("invalid menu id {%s}", menuItemId)
	}
	if err != nil {
		return err
	}

	return nil
}

func (mr *MenuRepository) GetById(menuItemId string) (*model.MenuItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var m model.MenuItem
	err := mr.col.FindOne(ctx, bson.D{{"id", menuItemId}}).Decode(&m)
	if err != nil {
		return &model.MenuItem{}, err
	}

	return &m, nil
}

func (mr *MenuRepository) GetAll() ([]model.MenuItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := mr.col.Find(ctx, bson.D{})
	if err != nil {
		return []model.MenuItem{}, err
	}

	var menus []model.MenuItem
	if err := cur.All(context.Background(), &menus); err != nil {
		return []model.MenuItem{}, err
	}

	return menus, nil
}
