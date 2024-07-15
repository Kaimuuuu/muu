package fiber

import (
	"fmt"
	"github.com/Kaimuuuu/muu/model"

	"github.com/gofiber/fiber/v2"
)

func chefRoleFilterer(c *fiber.Ctx) error {
	payload := c.Locals("payload").(JwtPayload)
	if payload.Role == model.Chef {
		return fmt.Errorf("chef is not allow to access this route")
	}

	return c.Next()
}
