package fiber

import (
	"kaimuu/model"
	"kaimuu/service/client"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func (f *FiberServer) AddAuthRoutes(clientTokenHandler func(*fiber.Ctx) error, employeeTokenHandler func(*fiber.Ctx) error) {
	routes := f.app.Group("/auth")

	routes.Post("/", func(c *fiber.Ctx) error {
		req := client.SignInRequest{}
		if err := c.BodyParser(&req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.Validate(req); err != nil {
			return f.errorHandler(c, err)
		}

		employee, err := f.employeeServ.SignIn(req.Email, req.Password)
		if err != nil {
			return f.errorHandler(c, err)
		}

		// Create the Claims
		claims := jwt.MapClaims{
			"employeeId": employee.Id,
			"role":       employee.Role,
			"exp":        time.Now().Add(time.Hour * f.config.JwtExpireHour).Unix(),
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(f.config.JwtSecret))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"token": t, "role": employee.Role})
	})

	routes.Get("/me/client", clientTokenHandler, func(c *fiber.Ctx) error {
		cli := c.Locals("client").(*model.Client)

		return c.Status(fiber.StatusOK).JSON(cli)
	})

	routes.Get("/me", employeeTokenHandler, toJwtPayloadHandler, func(c *fiber.Ctx) error {
		payload := c.Locals("payload").(JwtPayload)

		return c.Status(fiber.StatusOK).JSON(payload)
	})
}
