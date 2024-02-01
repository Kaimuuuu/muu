package fiber

import (
	"time"

	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

var TokenExpiredError = errors.New("token expired")
var InvalidToken = errors.New("token expired")

func (f *FiberServer) NewClientTokenHandler() func(*fiber.Ctx) error {
	return keyauth.New(keyauth.Config{
		Validator: func(c *fiber.Ctx, token string) (bool, error) {
			cli, err := f.tokenServ.Get(token)
			if err != nil {
				return false, InvalidToken
			}

			if time.Now().Compare(cli.Expire) >= 0 {
				return false, TokenExpiredError
			}

			c.Locals("client", cli)

			return true, nil
		},
		ErrorHandler: f.errorHandler,
	})
}
