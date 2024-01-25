package fiber

import (
	"fmt"
	"kaimuu/model"

	"github.com/gofiber/fiber/v2"
)

func waiterRoleFilterer(c *fiber.Ctx) error {
	payload := c.Locals("payload").(JwtPayload)
	if payload.Role == model.Waiter {
		return fmt.Errorf("waiter is not allow to access this route")
	}

	return c.Next()
}
