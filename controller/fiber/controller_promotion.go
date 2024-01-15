package fiber

import (
	"kaimuu/model"
	"kaimuu/service/promotion"

	"github.com/gofiber/fiber/v2"
)

func (f *FiberServer) AddPromotionRoutes(clientTokenHandler func(*fiber.Ctx) error, employeeTokenHandler func(*fiber.Ctx) error) {

	f.app.Get("/promotion/weight", clientTokenHandler, func(c *fiber.Ctx) error {
		cli := c.Locals("client").(*model.Client)

		weight, err := f.promotionServ.GetWeight(cli.PromotionId)
		if err != nil {
			c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"weight": weight})
	})

	routes := f.app.Group("/promotion", employeeTokenHandler, toJwtPayloadHandler)

	routes.Post("/", chefRoleFilterer, waiterRoleFilterer, func(c *fiber.Ctx) error {
		payload := c.Locals("payload").(JwtPayload)

		req := promotion.CreatePromotionRequest{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}

		ok, errMessage := f.Validate(req)
		if !ok {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(errMessage)
		}

		if err := f.promotionServ.CreatePromotion(req, payload.EmployeeId); err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}

		return nil
	})

	routes.Delete("/:promotionId", chefRoleFilterer, waiterRoleFilterer, func(c *fiber.Ctx) error {
		promotionId := c.Params("promotionId")

		if err := f.promotionServ.Delete(promotionId); err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}

		return nil
	})

	routes.Get("/", func(c *fiber.Ctx) error {
		promotions, err := f.promotionServ.GetPromotions()
		if err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(promotions)
	})

	routes.Get("/:promotionId/menu", func(c *fiber.Ctx) error {
		promotionId := c.Params("promotionId")

		promotionMenuItems, err := f.promotionServ.GetPromotionMenu(promotionId)
		if err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(promotionMenuItems)
	})

	routes.Put("/:promotionId", chefRoleFilterer, waiterRoleFilterer, func(c *fiber.Ctx) error {
		promotionId := c.Params("promotionId")

		req := promotion.UpdatePromotionRequest{}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}

		ok, errMessage := f.Validate(req)
		if !ok {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(errMessage)
		}

		if err := f.promotionServ.UpdatePromotion(promotionId, req); err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}

		return nil
	})
}
