package sql

import (
	"errors"

	"gorm.io/gorm"

	"github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/entities"
	internalErrors "github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/utils/errors"
)

type orderRepository struct {
	Conn *gorm.DB
}

func NewOrderRepository(Conn *gorm.DB) entities.IOrderRepository {
	return &orderRepository{Conn}
}

func (repo *orderRepository) GetOrderByID(ID string) (entities.Order, error) {
	var order entities.Order
	if err := repo.Conn.Preload("OrderItems").First(&order, "id = ?", ID).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return order, &internalErrors.ValidationError{Message: "Could not find order with provided ID"}
	} else {
		return order, err
	}
}

func (repo *orderRepository) CreateOrder(order entities.Order) (string, error) {
	err := repo.Conn.Create(&order).Error
	return order.ID, err
}
