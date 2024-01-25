package fiber

import (
	"fmt"
	"kaimuu/model"
	"kaimuu/service/client"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
)

func (f *FiberServer) AddClientRoutes(clientTokenHandler func(*fiber.Ctx) error, employeeTokenHandler func(*fiber.Ctx) error) {
	f.app.Get("/client/order", clientTokenHandler, func(c *fiber.Ctx) error {
		cli := c.Locals("client").(*model.Client)

		orders, err := f.clientServ.GetOrderHistory(cli)
		if err != nil {
			return f.errorHandler(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(orders)
	})

	f.app.Get("/qrcode/:token", func(c *fiber.Ctx) error {
		token := c.Params("token")

		content := fmt.Sprintf("%s/entrypoint?token=%s", os.Getenv("FRONTEND_URL"), token)
		qr, err := qrcode.Encode(content, 0, 256)
		if err != nil {
			return f.errorHandler(c, err)
		}

		return c.Type("png").Send(qr)
	})

	routes := f.app.Group("/client", employeeTokenHandler, toJwtPayloadHandler)

	routes.Post("/checkout/:token", chefRoleFilterer, func(c *fiber.Ctx) error {
		token := c.Params("token")

		if err := f.clientServ.Checkout(token); err != nil {
			return f.errorHandler(c, err)
		}

		return c.SendStatus(fiber.StatusCreated)
	})

	routes.Get("/checkout/:token", chefRoleFilterer, func(c *fiber.Ctx) error {
		token := c.Params("token")

		t, err := f.clientServ.CheckoutSummary(token)
		if err != nil {
			return f.errorHandler(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(t)
	})

	routes.Delete("/:token", chefRoleFilterer, waiterRoleFilterer, func(c *fiber.Ctx) error {
		token := c.Params("token")

		if err := f.clientServ.Delete(token); err != nil {
			return f.errorHandler(c, err)
		}

		return c.SendStatus(fiber.StatusOK)
	})

	routes.Post("/", chefRoleFilterer, func(c *fiber.Ctx) error {
		payload := c.Locals("payload").(JwtPayload)

		req := client.GenerateClientRequest{}
		if err := c.BodyParser(&req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.Validate(req); err != nil {
			return f.errorHandler(c, err)
		}

		token, err := f.clientServ.GenerateClient(req, payload.EmployeeId)
		if err != nil {
			return f.errorHandler(c, err)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"clientToken": token})
	})

	routes.Get("/", chefRoleFilterer, func(c *fiber.Ctx) error {
		clients, err := f.clientServ.GetClients()
		if err != nil {
			return f.errorHandler(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(clients)
	})
}
