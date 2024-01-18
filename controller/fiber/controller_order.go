package fiber

import (
	"kaimuu/model"
	"kaimuu/service/order"

	"github.com/gofiber/fiber/v2"
)

func (f *FiberServer) AddOrderRoutes(clientTokenHandler func(*fiber.Ctx) error, employeeTokenHandler func(*fiber.Ctx) error) {

	f.app.Post("/order", clientTokenHandler, func(c *fiber.Ctx) error {
		cli := c.Locals("client").(*model.Client)

		req := order.CreateOrderRequest{}
		if err := c.BodyParser(&req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.Validate(req); err != nil {
			return f.errorHandler(c, err)
		}

		if err := f.orderServ.CreateOrder(req, cli); err != nil {
			return f.errorHandler(c, err)
		}

		for _, roi := range req.OrderItems {
			f.srs.Increment(roi.MenuItemId, roi.Quantity)
		}

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
		orders, err := f.orderServ.GetPendingOrder()
		if err != nil {
			return f.errorHandler(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(orders)
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
}
