package sql

import (
	"errors"

	"gorm.io/gorm"

	"github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/entities"
	internalErrors "github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/utils/errors"
)

type pizzaRepository struct {
	Conn *gorm.DB
}

func NewPizzaRepository(Conn *gorm.DB) entities.IPizzaRepository {
	return &pizzaRepository{Conn}
}

func (repo *pizzaRepository) GetPizzas() ([]entities.Pizza, error) {
	var pizzas []entities.Pizza
	err := repo.Conn.Find(&pizzas).Error
	return pizzas, err
}

func (repo *pizzaRepository) GetPizzaByID(ID string) (entities.Pizza, error) {
	var pizza entities.Pizza
	if err := repo.Conn.First(&pizza, "id = ?", ID).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return pizza, &internalErrors.ValidationError{Message: "Could not find pizza with provided ID"}
	} else {
		return pizza, err
	}
}

func (repo *pizzaRepository) GetPizzaByName(name string) (entities.Pizza, error) {
	var pizza entities.Pizza
	if err := repo.Conn.Where("name = ?", name).First(&pizza).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return pizza, &internalErrors.ValidationError{Message: "Could not find pizza with provided name"}
	} else {
		return pizza, err
	}
}
