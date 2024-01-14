package employee

import (
	"kaimuu/model"

	"golang.org/x/crypto/bcrypt"
)

func (es *EmployeeService) SignIn(email, password string) (*model.Employee, error) {
	employee, err := es.employeeRepo.GetByEmail(email)
	if err != nil {
		return &model.Employee{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(password)); err != nil {
		return &model.Employee{}, err
	}

	return employee, nil
}
