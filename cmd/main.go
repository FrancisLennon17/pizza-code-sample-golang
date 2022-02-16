package main

import (
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/francislennon17/pizza-code-sample-golang/m/v2/docs"
	"github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/db"
	"github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/entities"
	orderUsecase "github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/order"
	orderHandler "github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/order/handler"
	orderSQL "github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/order/repository/sql"
	pizzaUsecase "github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/pizza"
	pizzaHandler "github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/pizza/handler"
	pizzaSQL "github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/pizza/repository/sql"
	"github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/utils"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// @title Pizza Code Sample
// @version 1.0
// @description This is a sample server golang pizza order service
// @termsOfService http://swagger.io/terms/
func main() {
	log.SetOutput(os.Stdout)

	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(middleware.Logger())

	db, err := db.GetConnection()
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&entities.Pizza{}, &entities.Order{}, &entities.OrderItem{}); err != nil { // bad practice
		panic(err)
	}

	pizzaRepo := pizzaSQL.NewPizzaRepository(db)
	pizzaUC := pizzaUsecase.NewPizzaUsecase(pizzaRepo)
	pizzaHandler.NewPizzaHandler(e, pizzaUC)

	orderRepo := orderSQL.NewOrderRepository(db)
	orderUC := orderUsecase.NewOrderUsecase(orderRepo, pizzaUC)
	orderHandler.NewOrderHandler(e, orderUC)

	e.Logger.Fatal(e.Start(":8080"))
}
