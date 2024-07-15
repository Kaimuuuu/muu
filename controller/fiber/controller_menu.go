package fiber

import (
	"github.com/Kaimuuuu/muu/model"
	"github.com/Kaimuuuu/muu/service/menu"
	"github.com/Kaimuuuu/muu/util/pubsub"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (f *FiberServer) AddMenuRoutes(clientTokenHandler func(*fiber.Ctx) error, employeeTokenHandler func(*fiber.Ctx) error) {
	ps := pubsub.New()

	f.app.Get("/menu", clientTokenHandler, func(c *fiber.Ctx) error {
		cli := c.Locals("client").(*model.Client)

		m, err := f.menuServ.Get(cli)
		if err != nil {
			return f.errorHandler(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(m)
	})

	f.app.Get("/menu/long-polling", clientTokenHandler, func(c *fiber.Ctx) error {
		cli := c.Locals("client").(*model.Client)

		ch, closeCh := ps.Subscribe()
		defer closeCh()

		to := make(chan struct{}, 1)

		go func() {
			time.Sleep(time.Hour * 2)
			to <- struct{}{}
			close(to)
		}()

		select {
		case <-ch:
			m, err := f.menuServ.Get(cli)
			if err != nil {
				return f.errorHandler(c, err)
			}
			return c.Status(fiber.StatusOK).JSON(m)
		case <-to:
			return c.SendStatus(fiber.StatusRequestTimeout)
		}
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

		ps.Publish()

		return c.SendStatus(fiber.StatusOK)
	})
}
