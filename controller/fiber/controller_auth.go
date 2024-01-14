package fiber

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type SignInRequest struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

func (f *FiberServer) AddAuthRoutes() {
	routes := f.app.Group("/auth")

	routes.Post("/sign-in", func(c *fiber.Ctx) error {
		req := SignInRequest{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}

		ok, errMessage := f.Validate(req)
		if !ok {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(errMessage)
		}

		employee, err := f.employeeServ.SignIn(req.Email, req.Password)
		if err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
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

		return c.JSON(fiber.Map{"token": t, "role": employee.Role})
	})
}
