package fiber

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

func (f *FiberServer) NewClientTokenHandler() func(*fiber.Ctx) error {
	return keyauth.New(keyauth.Config{
		Validator: func(c *fiber.Ctx, token string) (bool, error) {
			cli, err := f.tokenServ.Get(token)

			if err != nil {
				return false, keyauth.ErrMissingOrMalformedAPIKey
			}

			if time.Now().Compare(cli.Expire) >= 0 {
				if err := f.tokenServ.Delete(token); err != nil {
					return false, fmt.Errorf("cannot remove token")
				}
				return false, fmt.Errorf("token expired")
			}

			c.Locals("client", cli)

			return true, nil
		},
		ErrorHandler: f.errorHandler,
	})
}
