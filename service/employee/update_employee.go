package employee

import (
	"fmt"
	"kaimuu/model"
	"os"
)

func (es *EmployeeService) UpdateEmployee(employeeId string, req UpdateEmployeeRequest) error {
	empl, err := es.employeeRepo.GetById(employeeId)

	if err != nil {
		return err
	}

	empl.Name = req.Name
	empl.Age = req.Age
	empl.Role = model.EmployeeRole(req.Role)
	empl.ImagePath = req.ImagePath
	empl.Email = req.Email

	os.Remove(fmt.Sprintf("public/%s", empl.ImagePath))

	if err := es.employeeRepo.Update(employeeId, empl); err != nil {
		return err
	}

	return nil
}
