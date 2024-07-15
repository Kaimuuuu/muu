package fiber

import (
	"github.com/Kaimuuuu/muu/service/employee"

	"github.com/gofiber/fiber/v2"
)

func (f *FiberServer) AddEmployeeRoutes(employeeTokenHandler func(*fiber.Ctx) error) {
	routes := f.app.Group("/employee", employeeTokenHandler, toJwtPayloadHandler, chefRoleFilterer, waiterRoleFilterer)

	routes.Post("/", func(c *fiber.Ctx) error {
		payload := c.Locals("payload").(JwtPayload)

		req := employee.CreateEmployeeRequest{}
		if err := c.BodyParser(&req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.Validate(req); err != nil {
			return f.errorHandler(c, err)
		}

		genPass, err := f.employeeServ.Create(req, payload.EmployeeId)
		if err != nil {
			return f.errorHandler(c, err)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"generatedPassword": genPass})
	})

	routes.Delete("/:employeeId", func(c *fiber.Ctx) error {
		employeeId := c.Params("employeeId")

		if err := f.employeeServ.Delete(employeeId); err != nil {
			return f.errorHandler(c, err)
		}

		return c.SendStatus(fiber.StatusOK)
	})

	routes.Get("/", func(c *fiber.Ctx) error {
		employees, err := f.employeeServ.All()
		if err != nil {
			return f.errorHandler(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(employees)
	})

	routes.Put("/:employeeId", func(c *fiber.Ctx) error {
		employeeId := c.Params("employeeId")

		req := employee.UpdateEmployeeRequest{}
		if err := c.BodyParser(&req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.Validate(req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.employeeServ.Update(employeeId, req); err != nil {
			return f.errorHandler(c, err)
		}

		return c.SendStatus(fiber.StatusOK)
	})
}
