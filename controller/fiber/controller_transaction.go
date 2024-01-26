package fiber

import "github.com/gofiber/fiber/v2"

func (f *FiberServer) AddTransactionRoutes(clientTokenHandler func(*fiber.Ctx) error, employeeTokenHandler func(*fiber.Ctx) error) {
	routes := f.app.Group("/transaction", employeeTokenHandler, toJwtPayloadHandler)

	routes.Post("/checkout/:token", chefRoleFilterer, func(c *fiber.Ctx) error {
		token := c.Params("token")

		if err := f.transactionServ.Checkout(token); err != nil {
			return f.errorHandler(c, err)
		}

		return c.SendStatus(fiber.StatusCreated)
	})

	routes.Get("/checkout/:token", chefRoleFilterer, func(c *fiber.Ctx) error {
		token := c.Params("token")

		t, err := f.transactionServ.Summary(token)
		if err != nil {
			return f.errorHandler(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(t)
	})
}
