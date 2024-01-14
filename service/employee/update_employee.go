package employee

import "kaimuu/model"

func (es *EmployeeService) UpdateEmployee(employeeId string, req UpdateEmployeeRequest) error {
	empl, err := es.employeeRepo.GetById(employeeId)

	if err != nil {
		return err
	}

	empl.Name = req.Name
	empl.Age = req.Age
	empl.Role = model.EmployeeRole(req.Role)
	empl.Email = req.Email

	if err := es.employeeRepo.Update(employeeId, empl); err != nil {
		return err
	}

	return nil
}
