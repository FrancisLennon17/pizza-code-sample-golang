package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/entities"
	internalErrors "github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/utils/errors"
)

type OrderHandler struct {
	Usecase entities.IOrderUsecase
}

func NewOrderHandler(e *echo.Echo, uc entities.IOrderUsecase) {
	handler := &OrderHandler{uc}
	e.GET("api/orders/:id", handler.GetOrderByID)
	e.POST("api/orders", handler.CreateOrder)
}

// GET order by id
// @tags
// @Summary      Get order by ID
// @Accept       json
// @Produce      json
// @Param orderID path string true "Order ID"
// @Success      200  {object}  entities.Order
// @failure      400  {object}  errors.ValidationError
// @failure      500
// @Router       /api/orders/{orderID} [get]
func (x OrderHandler) GetOrderByID(ctx echo.Context) error {
	orderID := ctx.Param("id")
	if orderID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, &internalErrors.ValidationError{Message: "No order id provided"})
	}
	if order, err := x.Usecase.GetOrderByID(string(orderID)); err != nil {
		validationErr := &internalErrors.ValidationError{}
		if errors.As(err, &validationErr) {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		} else {
			log.Errorf("failed to fetch order with error: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	} else {
		return ctx.JSON(http.StatusOK, order)
	}
}

// POST order
// @tags
// @Summary      Creates an order
// @Accept       json
// @Produce      json
// @Param order body entities.Order true "Create order"
// @Success      200  {string}  orderID
// @failure      400  {object}  errors.ValidationError
// @failure      500
// @Router       /api/orders [post]
func (x OrderHandler) CreateOrder(ctx echo.Context) error {
	var order entities.Order

	if err := ctx.Bind(&order); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := ctx.Validate(order); err != nil {
		return err
	}
	if orderID, err := x.Usecase.CreateOrder(order); err != nil {
		validationError := &internalErrors.ValidationError{}
		if errors.As(err, &validationError) {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		} else {
			log.Errorf("failed to fetch order with error: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	} else {
		return ctx.JSON(http.StatusOK, orderID)
	}
}
