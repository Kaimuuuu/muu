package employee

import (
	"kaimuu/model"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (es *EmployeeService) Create(req CreateEmployeeRequest, employeeId string) (string, error) {
	_, err := es.employeeRepo.GetByEmail(req.Email)
	if err == nil {
		return "", EmailHaveBeenUsedError
	}

	genPass := uuid.NewString()

	hash, err := bcrypt.GenerateFromPassword([]byte(genPass), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	e := &model.Employee{
		Id:        uuid.NewString(),
		Name:      req.Name,
		Password:  string(hash),
		Age:       req.Age,
		Role:      model.EmployeeRole(req.Role),
		Email:     req.Email,
		ImagePath: req.ImagePath,
		CreatedAt: time.Now(),
		CreatedBy: employeeId,
	}

	if err := es.employeeRepo.Insert(e); err != nil {
		return "", err
	}

	return genPass, nil
}
