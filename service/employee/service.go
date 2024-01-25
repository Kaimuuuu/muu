package employee

import (
	"kaimuu/model"

	"github.com/cockroachdb/errors"
)

var (
	InvalidEmailError      = errors.New("invalid email")
	InvalidPasswordError   = errors.New("invalid password")
	EmailHaveBeenUsedError = errors.New("email have been used")
)

func NewEmployeeService(employeeRepo EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		employeeRepo: employeeRepo,
	}
}

type EmployeeService struct {
	employeeRepo EmployeeRepository
}

type EmployeeRepository interface {
	Insert(e *model.Employee) error
	Update(id string, empl *model.Employee) error
	Delete(id string) error
	GetById(id string) (*model.Employee, error)
	GetAll() ([]model.Employee, error)
	GetByEmail(email string) (*model.Employee, error)
}

type CreateEmployeeRequest struct {
	Name      string `validate:"required" json:"name"`
	Age       int8   `validate:"required,number" json:"age"`
	Role      int8   `validate:"min=0" json:"role"`
	ImagePath string `validate:"" json:"imagePath"`
	Email     string `validate:"required,email" json:"email"`
}

type UpdateEmployeeRequest struct {
	Name      string `validate:"required" json:"name"`
	Age       int8   `validate:"required,number" json:"age"`
	Role      int8   `validate:"min=0" json:"role"`
	ImagePath string `validate:"" json:"imagePath"`
	Email     string `validate:"required,email" json:"email"`
}
