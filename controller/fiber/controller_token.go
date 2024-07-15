package fiber

import (
	"fmt"
	"github.com/Kaimuuuu/muu/service/token"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
)

func (f *FiberServer) AddTokenRoutes(clientTokenHandler func(*fiber.Ctx) error, employeeTokenHandler func(*fiber.Ctx) error) {
	f.app.Get("/qrcode/:token", func(c *fiber.Ctx) error {
		token := c.Params("token")

		content := fmt.Sprintf("%s/entrypoint?token=%s", os.Getenv("FRONTEND_URL"), token)
		qr, err := qrcode.Encode(content, 0, 256)
		if err != nil {
			return f.errorHandler(c, err)
		}

		return c.Type("png").Send(qr)
	})

	routes := f.app.Group("/token", employeeTokenHandler, toJwtPayloadHandler)

	routes.Delete("/:token", chefRoleFilterer, waiterRoleFilterer, func(c *fiber.Ctx) error {
		token := c.Params("token")

		if err := f.tokenServ.Delete(token); err != nil {
			return f.errorHandler(c, err)
		}

		return c.SendStatus(fiber.StatusOK)
	})

	routes.Post("/", chefRoleFilterer, func(c *fiber.Ctx) error {
		payload := c.Locals("payload").(JwtPayload)

		req := token.GenerateTokenRequest{}
		if err := c.BodyParser(&req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.Validate(req); err != nil {
			return f.errorHandler(c, err)
		}

		token, err := f.tokenServ.Generate(req, payload.EmployeeId)
		if err != nil {
			return f.errorHandler(c, err)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"clientToken": token})
	})

	routes.Get("/", chefRoleFilterer, func(c *fiber.Ctx) error {
		clients, err := f.tokenServ.All()
		if err != nil {
			return f.errorHandler(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(clients)
	})
}
