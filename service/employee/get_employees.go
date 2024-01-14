package employee

import "kaimuu/model"

func (es *EmployeeService) GetEmployees() ([]model.Employee, error) {
	employees, err := es.employeeRepo.GetAll()

	if err != nil {
		return []model.Employee{}, err
	}

	return employees, nil
}
