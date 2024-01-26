package employee

import (
	"kaimuu/model"

	"golang.org/x/crypto/bcrypt"
)

func (es *EmployeeService) SignIn(email, password string) (*model.Employee, error) {
	e, err := es.employeeRepo.GetByEmail(email)
	if err != nil {
		return &model.Employee{}, InvalidEmailError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(e.Password), []byte(password)); err != nil {
		return &model.Employee{}, InvalidPasswordError
	}

	return e, nil
}
