package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID         string      `json:"id" gorm:"primaryKey"`
	Name       string      `json:"name" validate:"required" example:"Jimmy"`
	OrderItems []OrderItem `json:"orderItems" validate:"required,dive"`
	CreatedAt  time.Time   `json:"-" gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `json:"-" gorm:"autoUpdateTime"`
}

// overwriting Gorm's BeforeCreate function to auto create a UUID
func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New().String()
	return
}

type OrderItem struct {
	ID        string `json:"-" gorm:"primaryKey"`
	Size      string `json:"size" validate:"required,oneof=Large Medium Small" example:"Large"`
	OrderID   string `json:"-"`
	PizzaID   string `json:"-"`
	PizzaName string `json:"pizzaName" gorm:"-" validate:"required" example:"Meatfeast"`
}

// overwriting Gorm's BeforeCreate function to auto create a UUID
func (i *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = uuid.New().String()
	return
}

//go:generate mockgen -destination=./mocks/mock_order_usecase.go -package=mocks github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/entities IOrderUsecase
type IOrderUsecase interface {
	GetOrderByID(string) (Order, error)
	CreateOrder(Order) (string, error)
}

//go:generate mockgen -destination=./mocks/mock_order_repository.go -package=mocks github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/entities IOrderRepository
type IOrderRepository interface {
	GetOrderByID(string) (Order, error)
	CreateOrder(Order) (string, error)
}
