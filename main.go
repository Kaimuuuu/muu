package main

import (
	"context"
	"kaimuu/controller/fiber"
	"kaimuu/repository"
	"kaimuu/service/employee"
	"kaimuu/service/menu"
	"kaimuu/service/order"
	"kaimuu/service/promotion"
	srs "kaimuu/service/simple-recommandation-system"
	"kaimuu/service/token"
	"kaimuu/service/transaction"
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
	promotionRepo := repository.NewPromotionRepository(db)
	tokenRepo := repository.NewTokenRepository(db)

	srs := srs.New(menuRepo, promotionRepo)

	employeeServ := employee.NewEmployeeService(employeeRepo)
	promotoinServ := promotion.NewPromotionService(promotionRepo, menuRepo, tokenRepo)
	orderServ := order.NewOrderService(orderRepo, menuRepo, promotionRepo, tokenRepo)
	menuServ := menu.NewMenuService(menuRepo, orderServ, promotionRepo, promotoinServ)
	tokenServ := token.NewTokenService(tokenRepo, promotionRepo)
	transactionServ := transaction.NewTransactionService(transactionRepo, orderRepo, promotionRepo, tokenRepo)

	cfg := fiber.FiberConfig{
		Port:          os.Getenv("PORT"),
		JwtSecret:     os.Getenv("JWT_SECRET"),
		JwtExpireHour: 10,
	}

	fiber := fiber.New(cfg, transactionServ, tokenServ, employeeServ, menuServ, orderServ, promotoinServ, srs)
	fiber.Start()
}
