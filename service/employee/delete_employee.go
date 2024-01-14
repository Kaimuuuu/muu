package employee

func (es *EmployeeService) Delete(employeeId string) error {
	if err := es.employeeRepo.Delete(employeeId); err != nil {
		return err
	}

	return nil
}
