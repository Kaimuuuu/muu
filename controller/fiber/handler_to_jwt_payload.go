package fiber

import (
	"kaimuu/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func toJwtPayloadHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	role := int8(claims["role"].(float64))
	payload := JwtPayload{
		EmployeeId: claims["employeeId"].(string),
		Role:       model.EmployeeRole(role),
	}
	c.Locals("payload", payload)

	return c.Next()
}
