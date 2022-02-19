package order

import (
	log "github.com/sirupsen/logrus"

	"github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/entities"
)

type orderUsecase struct {
	repo    entities.IOrderRepository
	pizzaUc entities.IPizzaUsecase
}

func NewOrderUsecase(repo entities.IOrderRepository, pizzaUc entities.IPizzaUsecase) entities.IOrderUsecase {
	return &orderUsecase{
		repo:    repo,
		pizzaUc: pizzaUc,
	}
}

func (x *orderUsecase) GetOrderByID(ID string) (entities.Order, error) {
	if order, err := x.repo.GetOrderByID(ID); err != nil {
		return entities.Order{}, err
	} else {
		for i, orderItem := range order.OrderItems {
			if pizza, pizzaErr := x.pizzaUc.GetPizzaByID(orderItem.PizzaID); pizzaErr != nil {
				log.Errorf("Get Order By ID - failed to fetch pizza name by ID with error: %v", err)
				return entities.Order{}, pizzaErr
			} else {
				order.OrderItems[i].PizzaName = pizza.Name
			}
		}
		return order, nil
	}
}

func (x *orderUsecase) CreateOrder(order entities.Order) (string, error) {
	for i, orderItem := range order.OrderItems {
		if pizza, err := x.pizzaUc.GetPizzaByName(orderItem.PizzaName); err != nil {
			log.Errorf("Create Order - failed to fetch pizza name by ID with error: %v", err)
			return pizza.ID, err
		} else {
			order.OrderItems[i].PizzaID = pizza.ID
		}
	}
	return x.repo.CreateOrder(order)
}
