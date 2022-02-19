package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/entities"
)

type PizzaHandler struct {
	Usecase entities.IPizzaUsecase
}

func NewPizzaHandler(e *echo.Echo, uc entities.IPizzaUsecase) {
	handler := &PizzaHandler{
		Usecase: uc,
	}
	e.GET("api/pizzas", handler.GetPizzas)
}

// GET pizzas
// @tags
// @Summary      Get list of pizzas
// @Description  Fetches the list of pizzas that can be ordered
// @Accept       json
// @Produce      json
// @Success      200  {array}  entities.Pizza
// @failure      500
// @Router       /api/pizzas [get]
func (x PizzaHandler) GetPizzas(ctx echo.Context) error {
	if pizzas, err := x.Usecase.GetPizzas(); err != nil {
		log.Errorf("failed to fetch pizzas with error: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	} else {
		return ctx.JSON(http.StatusOK, pizzas)
	}
}
