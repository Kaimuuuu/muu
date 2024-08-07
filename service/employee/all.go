package employee

import "github.com/Kaimuuuu/muu/model"

func (es *EmployeeService) All() ([]model.Employee, error) {
	employees, err := es.employeeRepo.All()

	if err != nil {
		return []model.Employee{}, err
	}

	return employees, nil
}
