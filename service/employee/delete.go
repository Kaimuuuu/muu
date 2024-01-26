package employee

import (
	"fmt"
	"os"
)

func (es *EmployeeService) Delete(employeeId string) error {
	e, err := es.employeeRepo.GetById(employeeId)
	if err != nil {
		return err
	}

	os.Remove(fmt.Sprintf("public/%s", e.ImagePath))

	if err := es.employeeRepo.Delete(employeeId); err != nil {
		return err
	}

	return nil
}
