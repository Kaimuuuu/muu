package fiber

import (
	"github.com/Kaimuuuu/muu/service/employee"
	"github.com/Kaimuuuu/muu/service/order"
	"github.com/Kaimuuuu/muu/service/promotion"
	"github.com/Kaimuuuu/muu/service/token"

	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

type ErrorInfo struct {
	StatusCode int
	Message    string
}

type ErrorResponse struct {
	ErrMessage string `json:"errMessage"`
}

var errorList = map[error]ErrorInfo{
	employee.InvalidEmailError:      {400, "email ไม่ถูกต้อง"},
	employee.InvalidPasswordError:   {400, "รหัสผ่านไม่ถูกต้อง"},
	employee.EmailHaveBeenUsedError: {400, "email นี้ถูกใช้งานแล้ว"},
	order.OrderInvalidMenuItemError: {400, "menu item id ไม่ถูกต้อง"},
	order.MenuItemOutOfStockError:   {400, "มีเมนูที่สถาณะเป็น \"หมด\""},
	order.WeightExceededError:       {400, "น้ำหนักต่อคำสั่งอาหารเกินกำหมด"},
	promotion.ClientInUsedError:     {400, "ไม่สามารถแก้ไขได้เนื่องจากมีลูกค้าใช้งานโปรโมชั่นนี้อยู่"},
	token.TableInUsedError:          {400, "โต๊ะนี้ถูกใช้งานแล้ว"},
	TokenExpiredError:               {500, "token นี้หมดอายุการใช้งานแล้ว"},
	InvalidToken:                    {500, "token ไม่ถูกต้อง"},
	order.InvalidOrderQuantity:      {400, "สั่งอาหารเกินปริมานที่กำหนด"},
}

func (f *FiberServer) sendError(c *fiber.Ctx, errInfo ErrorInfo) error {
	return c.Status(errInfo.StatusCode).JSON(ErrorResponse{
		ErrMessage: errInfo.Message,
	})
}

func (f *FiberServer) errorHandler(c *fiber.Ctx, err error) error {
	for iErr, info := range errorList {
		if errors.Is(err, iErr) {
			return f.sendError(c, info)
		}
	}

	return f.sendError(c, ErrorInfo{StatusCode: 500, Message: err.Error()})
}
