package fiber

import (
	"kaimuu/model"
	"kaimuu/service/menu"

	"github.com/gofiber/fiber/v2"
)

func (f *FiberServer) AddMenuRoutes(clientTokenHandler func(*fiber.Ctx) error, employeeTokenHandler func(*fiber.Ctx) error) {
	f.app.Get("/menu", clientTokenHandler, func(c *fiber.Ctx) error {
		cli := c.Locals("client").(*model.Client)

		m, err := f.menuServ.Get(cli)
		if err != nil {
			return f.errorHandler(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(m)
	})

	f.app.Get("/menu/recommand", clientTokenHandler, func(c *fiber.Ctx) error {
		cli := c.Locals("client").(*model.Client)

		m, err := f.srs.Get(cli.PromotionId)
		if err != nil {
			return f.errorHandler(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(m)
	})

	routes := f.app.Group("/menu", employeeTokenHandler, toJwtPayloadHandler)

	routes.Get("/edit", func(c *fiber.Ctx) error {
		m, err := f.menuServ.All()
		if err != nil {
			return f.errorHandler(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(m)
	})

	routes.Post("/", chefRoleFilterer, waiterRoleFilterer, func(c *fiber.Ctx) error {
		payload := c.Locals("payload").(JwtPayload)

		req := menu.CreateMenuRequest{}
		if err := c.BodyParser(&req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.Validate(req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.menuServ.Create(req, payload.EmployeeId); err != nil {
			return f.errorHandler(c, err)
		}

		return c.SendStatus(fiber.StatusCreated)
	})

	routes.Delete("/:menuId", chefRoleFilterer, waiterRoleFilterer, func(c *fiber.Ctx) error {
		menuId := c.Params("menuId")

		if err := f.menuServ.Delete(menuId); err != nil {
			return f.errorHandler(c, err)
		}

		return c.SendStatus(fiber.StatusOK)
	})

	routes.Put("/:menuId", chefRoleFilterer, waiterRoleFilterer, employeeTokenHandler, func(c *fiber.Ctx) error {
		menuId := c.Params("menuId")

		req := menu.UpdateMenuRequest{}
		if err := c.BodyParser(&req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.Validate(req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.menuServ.Update(menuId, req); err != nil {
			return f.errorHandler(c, err)
		}

		return c.SendStatus(fiber.StatusOK)
	})

	routes.Put("/:menuId/out-of-stock", waiterRoleFilterer, func(c *fiber.Ctx) error {
		menuId := c.Params("menuId")

		req := menu.OutOfStockRequest{}
		if err := c.BodyParser(&req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.Validate(req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.menuServ.UpdateOutOfStock(menuId, req.IsOutOfStock); err != nil {
			return f.errorHandler(c, err)
		}

		return c.SendStatus(fiber.StatusOK)
	})
}
