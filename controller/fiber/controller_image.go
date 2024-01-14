package fiber

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (f *FiberServer) AddImageRoutes(employeeTokenHandler func(*fiber.Ctx) error) {
	routes := f.app.Group("/image", employeeTokenHandler, toJwtPayloadHandler, chefRoleFilterer, waiterRoleFilterer)

	routes.Post("/", func(c *fiber.Ctx) error {
		imageFile, err := c.FormFile("image")
		if err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}
		if err := c.SaveFile(imageFile, fmt.Sprintf("./public/%s", imageFile.Filename)); err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}

		return c.JSON(fiber.Map{"imagePath": imageFile.Filename})
	})
}
