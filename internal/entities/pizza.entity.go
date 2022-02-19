package entities

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Pizza struct {
	ID          string         `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" example:"Meatfeast"`
	Price       float64        `json:"price" example:"14.50"`
	Ingredients pq.StringArray `json:"ingredients" gorm:"type:text[]" swaggertype:"array,string" example:"Meat,Cheese,Sauce"`
}

// overwriting Gorm's BeforeCreate function to auto create a UUID
func (p *Pizza) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return
}

//go:generate mockgen -destination=./mocks/mock_pizza_usecase.go -package=mocks github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/entities IPizzaUsecase
type IPizzaUsecase interface {
	GetPizzas() ([]Pizza, error)
	GetPizzaByName(string) (Pizza, error)
	GetPizzaByID(string) (Pizza, error)
}

type IPizzaRepository interface {
	GetPizzas() ([]Pizza, error)
	GetPizzaByName(string) (Pizza, error)
	GetPizzaByID(string) (Pizza, error)
}
