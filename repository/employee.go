package repository

import (
	"context"
	"kaimuu/model"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepository struct {
	col *mongo.Collection
}

func NewEmployeeRepository(db *mongo.Database) *EmployeeRepository {
	return &EmployeeRepository{
		col: db.Collection("employees"),
	}
}

func (er *EmployeeRepository) Insert(empl *model.Employee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := er.col.InsertOne(ctx, empl)
	if err != nil {
		return err
	}

	return nil
}

func (er *EmployeeRepository) Update(employeeId string, empl *model.Employee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := er.col.UpdateOne(ctx, bson.D{{"id", employeeId}}, bson.D{{"$set", empl}})
	if result.MatchedCount == 0 {
		return errors.Errorf("invalid employee id '%s'", employeeId)
	}
	if err != nil {
		return err
	}

	return nil
}

func (er *EmployeeRepository) Delete(employeeId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := er.col.DeleteOne(ctx, bson.D{{"id", employeeId}})
	if result.DeletedCount == 0 {
		return errors.Errorf("invalid employee id '%s'", employeeId)
	}
	if err != nil {
		return err
	}

	return nil
}

func (er *EmployeeRepository) GetById(employeeId string) (*model.Employee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var empl model.Employee
	err := er.col.FindOne(ctx, bson.D{{"id", employeeId}}).Decode(&empl)
	if err != nil {
		return &model.Employee{}, err
	}

	return &empl, nil
}

func (er *EmployeeRepository) GetAll() ([]model.Employee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := er.col.Find(ctx, bson.D{})
	if err != nil {
		return []model.Employee{}, err
	}

	var employees []model.Employee
	if err := cur.All(context.Background(), &employees); err != nil {
		return []model.Employee{}, err
	}

	return employees, nil
}

func (er *EmployeeRepository) GetByEmail(email string) (*model.Employee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var empl model.Employee
	err := er.col.FindOne(ctx, bson.D{{"email", email}}).Decode(&empl)
	if err != nil {
		return &model.Employee{}, err
	}

	return &empl, nil
}
