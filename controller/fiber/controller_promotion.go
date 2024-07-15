package fiber

import (
	"github.com/Kaimuuuu/muu/service/promotion"

	"github.com/gofiber/fiber/v2"
)

func (f *FiberServer) AddPromotionRoutes(clientTokenHandler func(*fiber.Ctx) error, employeeTokenHandler func(*fiber.Ctx) error) {
	routes := f.app.Group("/promotion", employeeTokenHandler, toJwtPayloadHandler)

	routes.Post("/", chefRoleFilterer, waiterRoleFilterer, func(c *fiber.Ctx) error {
		payload := c.Locals("payload").(JwtPayload)

		req := promotion.CreatePromotionRequest{}
		if err := c.BodyParser(&req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.Validate(req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.promotionServ.Create(req, payload.EmployeeId); err != nil {
			return f.errorHandler(c, err)
		}

		return c.SendStatus(fiber.StatusCreated)
	})

	routes.Delete("/:promotionId", chefRoleFilterer, waiterRoleFilterer, func(c *fiber.Ctx) error {
		promotionId := c.Params("promotionId")
		if err := f.promotionServ.Delete(promotionId); err != nil {
			return f.errorHandler(c, err)
		}

		return c.SendStatus(fiber.StatusOK)
	})

	routes.Get("/", func(c *fiber.Ctx) error {
		promotions, err := f.promotionServ.All()
		if err != nil {
			return f.errorHandler(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(promotions)
	})

	routes.Get("/:promotionId/menu", func(c *fiber.Ctx) error {
		promotionId := c.Params("promotionId")

		promotionMenuItems, err := f.promotionServ.GetMenu(promotionId)
		if err != nil {
			return f.errorHandler(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(promotionMenuItems)
	})

	routes.Put("/:promotionId", chefRoleFilterer, waiterRoleFilterer, func(c *fiber.Ctx) error {
		promotionId := c.Params("promotionId")

		req := promotion.UpdatePromotionRequest{}
		if err := c.BodyParser(&req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.Validate(req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.promotionServ.Update(promotionId, req); err != nil {
			return f.errorHandler(c, err)
		}

		return c.SendStatus(fiber.StatusOK)
	})
}
