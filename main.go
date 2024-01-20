package main

import (
	"context"
	"kaimuu/controller/fiber"
	"kaimuu/repository"
	cli "kaimuu/service/client"
	"kaimuu/service/employee"
	token "kaimuu/service/in_memory_token"
	"kaimuu/service/menu"
	"kaimuu/service/order"
	"kaimuu/service/promotion"
	simplerecommandationsystem "kaimuu/service/simple-recommandation-system"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	db := client.Database(os.Getenv("MONGO_DB"))

	menuRepo := repository.NewMenuRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)
	promotoinRepo := repository.NewPromotionRepository(db)

	srs := simplerecommandationsystem.New(menuRepo, promotoinRepo)

	tokenStorage := token.NewInMemoryTokenStorage()
	employeeServ := employee.NewEmployeeService(employeeRepo)
	promotoinServ := promotion.NewPromotionService(promotoinRepo, menuRepo, tokenStorage)
	orderServ := order.NewOrderService(orderRepo, menuRepo, promotoinServ, tokenStorage)
	menuServ := menu.NewMenuService(menuRepo, promotoinServ, orderServ, promotoinRepo)
	clientServ := cli.NewClientService(transactionRepo, orderServ, tokenStorage, promotoinServ)

	cfg := fiber.FiberConfig{
		Port:          os.Getenv("PORT"),
		JwtSecret:     os.Getenv("JWT_SECRET"),
		JwtExpireHour: 10,
	}

	fiber := fiber.New(cfg, clientServ, employeeServ, menuServ, orderServ, promotoinServ, tokenStorage, srs)
	fiber.Start()
}
