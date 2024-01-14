package fiber

import (
	"kaimuu/service/employee"

	"github.com/gofiber/fiber/v2"
)

func (f *FiberServer) AddEmployeeRoutes(employeeTokenHandler func(*fiber.Ctx) error) {
	routes := f.app.Group("/employee", employeeTokenHandler, toJwtPayloadHandler, chefRoleFilterer, waiterRoleFilterer)

	routes.Post("/", func(c *fiber.Ctx) error {
		payload := c.Locals("payload").(JwtPayload)

		req := employee.CreateEmployeeRequest{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}

		ok, errMessage := f.Validate(req)
		if !ok {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(errMessage)
		}

		genPass, err := f.employeeServ.CreateEmployee(req, payload.EmployeeId)
		if err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}

		return c.SendString(genPass)
	})

	routes.Delete("/:employeeId", func(c *fiber.Ctx) error {
		employeeId := c.Params("employeeId")

		if err := f.employeeServ.Delete(employeeId); err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}

		return nil
	})

	routes.Get("/", func(c *fiber.Ctx) error {
		employees, err := f.employeeServ.GetEmployees()
		if err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(employees)
	})

	routes.Put("/:employeeId", func(c *fiber.Ctx) error {
		employeeId := c.Params("employeeId")

		req := employee.UpdateEmployeeRequest{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}

		ok, errMessage := f.Validate(req)
		if !ok {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(errMessage)
		}

		if err := f.employeeServ.UpdateEmployee(employeeId, req); err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}

		return nil
	})
}
