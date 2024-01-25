package fiber

import (
	"fmt"
	"kaimuu/model"
	"kaimuu/service/client"
	"kaimuu/service/employee"
	"kaimuu/service/menu"
	"kaimuu/service/order"
	"kaimuu/service/promotion"
	simplerecommandationsystem "kaimuu/service/simple-recommandation-system"
	"os"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberServer struct {
	app           *fiber.App
	validator     *validator.Validate
	config        FiberConfig
	clientServ    *client.ClientService
	employeeServ  *employee.EmployeeService
	menuServ      *menu.MenuService
	orderServ     *order.OrderService
	promotionServ *promotion.PromotionService
	tokenStorage  TokenStorage
	srs           *simplerecommandationsystem.SimpleRecommandationSystem
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

func New(config FiberConfig, clientServ *client.ClientService, employeeServ *employee.EmployeeService, menuServ *menu.MenuService, orderServ *order.OrderService, promotionServ *promotion.PromotionService, tokenStorage TokenStorage, srs *simplerecommandationsystem.SimpleRecommandationSystem) *FiberServer {
	return &FiberServer{
		app:           fiber.New(),
		validator:     validator.New(validator.WithRequiredStructEnabled()),
		config:        config,
		clientServ:    clientServ,
		employeeServ:  employeeServ,
		menuServ:      menuServ,
		orderServ:     orderServ,
		promotionServ: promotionServ,
		tokenStorage:  tokenStorage,
		srs:           srs,
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
	f.AddClientRoutes(clientTokenHandler, employeeTokenHandler)
	f.AddEmployeeRoutes(employeeTokenHandler)
	f.AddPromotionRoutes(clientTokenHandler, employeeTokenHandler)
	f.AddImageRoutes(employeeTokenHandler)
	f.AddAuthRoutes(clientTokenHandler, employeeTokenHandler)

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
