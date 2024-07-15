package fiber

import (
	"fmt"
	"github.com/Kaimuuuu/muu/model"
	"github.com/Kaimuuuu/muu/service/employee"
	"github.com/Kaimuuuu/muu/service/menu"
	"github.com/Kaimuuuu/muu/service/order"
	"github.com/Kaimuuuu/muu/service/promotion"
	srs "github.com/Kaimuuuu/muu/service/simple-recommandation-system"
	"github.com/Kaimuuuu/muu/service/token"
	"github.com/Kaimuuuu/muu/service/transaction"
	"os"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberServer struct {
	app             *fiber.App
	validator       *validator.Validate
	config          FiberConfig
	transactionServ *transaction.TransactionService
	tokenServ       *token.TokenService
	employeeServ    *employee.EmployeeService
	menuServ        *menu.MenuService
	orderServ       *order.OrderService
	promotionServ   *promotion.PromotionService
	srs             *srs.SimpleRecommandationSystem
}

type TokenStorage interface {
	Get(key string) (*model.Client, error)
	Set(key string, value *model.Client) error
	Remove(key string) error
	GetAll() ([]model.Client, error)
}

type FiberConfig struct {
	Port          string
	JwtSecret     string
	JwtExpireHour time.Duration
}

type JwtPayload struct {
	EmployeeId string
	Role       model.EmployeeRole
}

func New(config FiberConfig, transactionServ *transaction.TransactionService, tokenServ *token.TokenService, employeeServ *employee.EmployeeService, menuServ *menu.MenuService, orderServ *order.OrderService, promotionServ *promotion.PromotionService, srs *srs.SimpleRecommandationSystem) *FiberServer {
	return &FiberServer{
		app:             fiber.New(),
		validator:       validator.New(validator.WithRequiredStructEnabled()),
		config:          config,
		transactionServ: transactionServ,
		tokenServ:       tokenServ,
		employeeServ:    employeeServ,
		menuServ:        menuServ,
		orderServ:       orderServ,
		promotionServ:   promotionServ,
		srs:             srs,
	}
}

func (f *FiberServer) Start() {
	f.app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins:     os.Getenv("FRONTEND_URL"),
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	f.app.Static("/", "./public")

	clientTokenHandler := f.NewClientTokenHandler()
	employeeTokenHandler := f.NewEmployeeTokenHandler()

	f.AddMenuRoutes(clientTokenHandler, employeeTokenHandler)
	f.AddOrderRoutes(clientTokenHandler, employeeTokenHandler)
	f.AddTokenRoutes(clientTokenHandler, employeeTokenHandler)
	f.AddEmployeeRoutes(employeeTokenHandler)
	f.AddPromotionRoutes(clientTokenHandler, employeeTokenHandler)
	f.AddImageRoutes(employeeTokenHandler)
	f.AddAuthRoutes(clientTokenHandler, employeeTokenHandler)
	f.AddTransactionRoutes(clientTokenHandler, employeeTokenHandler)

	f.app.Listen(":" + f.config.Port)
}

func (f *FiberServer) Validate(req interface{}) error {
	errs := f.validator.Struct(req)
	if errs != nil {
		errMessage := make([]string, 0)
		for _, err := range errs.(validator.ValidationErrors) {
			errMessage = append(errMessage, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.Field(),
				err.Value(),
				err.Tag(),
			))
		}
		return errors.New(strings.Join(errMessage, " and "))
	}

	return nil
}
