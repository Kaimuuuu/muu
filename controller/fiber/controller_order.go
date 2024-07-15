package fiber

import (
	"github.com/Kaimuuuu/muu/model"
	"github.com/Kaimuuuu/muu/service/order"
	"github.com/Kaimuuuu/muu/util/pubsub"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (f *FiberServer) AddOrderRoutes(clientTokenHandler func(*fiber.Ctx) error, employeeTokenHandler func(*fiber.Ctx) error) {
	ps := pubsub.New()

	f.app.Get("/order/client", clientTokenHandler, func(c *fiber.Ctx) error {
		cli := c.Locals("client").(*model.Client)

		orders, err := f.orderServ.GetOrderByToken(cli.Token)
		if err != nil {
			return f.errorHandler(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(orders)
	})

	f.app.Post("/order", clientTokenHandler, func(c *fiber.Ctx) error {
		cli := c.Locals("client").(*model.Client)

		req := order.CreateOrderRequest{}
		if err := c.BodyParser(&req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.Validate(req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.orderServ.Create(req, cli); err != nil {
			return f.errorHandler(c, err)
		}

		for _, roi := range req.OrderItems {
			f.srs.Increment(roi.MenuItemId, roi.Quantity)
		}

		ps.Publish()

		return c.SendStatus(fiber.StatusCreated)
	})

	routes := f.app.Group("/order", employeeTokenHandler, toJwtPayloadHandler)

	routes.Delete("/:orderId", chefRoleFilterer, waiterRoleFilterer, func(c *fiber.Ctx) error {
		orderId := c.Params("orderId")

		if err := f.orderServ.Delete(orderId); err != nil {
			return f.errorHandler(c, err)
		}

		return c.SendStatus(fiber.StatusOK)
	})

	routes.Get("/pending", func(c *fiber.Ctx) error {
		orders, err := f.orderServ.GetPendingOrders()
		if err != nil {
			return f.errorHandler(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(orders)
	})

	routes.Get("/pending/long-polling", func(c *fiber.Ctx) error {
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
			orders, err := f.orderServ.GetPendingOrders()
			if err != nil {
				return f.errorHandler(c, err)
			}

			return c.Status(fiber.StatusOK).JSON(orders)
		case <-to:
			return c.SendStatus(fiber.StatusRequestTimeout)
		}
	})

	routes.Put("/status/:orderId", waiterRoleFilterer, func(c *fiber.Ctx) error {
		orderId := c.Params("orderId")

		req := order.UpdateOrderStatusRequest{}
		if err := c.BodyParser(&req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.Validate(req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.orderServ.UpdateOrderStatus(req, orderId); err != nil {
			return f.errorHandler(c, err)
		}

		return c.SendStatus(fiber.StatusOK)
	})

	routes.Put("/status/:orderId/items", waiterRoleFilterer, func(c *fiber.Ctx) error {
		orderId := c.Params("orderId")

		req := order.UpdateOrderItemsStatusRequest{}
		if err := c.BodyParser(&req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.Validate(req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.orderServ.UpdateOrderItemStatus(req, orderId); err != nil {
			return f.errorHandler(c, err)
		}

		return c.SendStatus(fiber.StatusOK)
	})
}
