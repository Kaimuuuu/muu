package fiber

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (f *FiberServer) AddImageRoutes(employeeTokenHandler func(*fiber.Ctx) error) {
	routes := f.app.Group("/image", employeeTokenHandler, toJwtPayloadHandler, chefRoleFilterer, waiterRoleFilterer)

	routes.Post("/", func(c *fiber.Ctx) error {
		file, err := c.FormFile("image")
		if err != nil {
			return f.errorHandler(c, err)
		}
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		if err := c.SaveFile(file, fmt.Sprintf("./public/%s", filename)); err != nil {
			return f.errorHandler(c, err)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"imagePath": filename})
	})
}
