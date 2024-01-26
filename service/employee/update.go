package employee

import (
	"fmt"
	"kaimuu/model"
	"os"
)

func (es *EmployeeService) Update(employeeId string, req UpdateEmployeeRequest) error {
	e, err := es.employeeRepo.GetById(employeeId)

	if err != nil {
		return err
	}

	os.Remove(fmt.Sprintf("public/%s", e.ImagePath))

	e.Name = req.Name
	e.Age = req.Age
	e.Role = model.EmployeeRole(req.Role)
	e.ImagePath = req.ImagePath
	e.Email = req.Email

	if err := es.employeeRepo.Update(employeeId, e); err != nil {
		return err
	}

	return nil
}
