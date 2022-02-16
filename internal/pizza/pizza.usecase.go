package pizza

import (
	"github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/entities"
)

type pizzaUsecase struct {
	repo entities.IPizzaRepository
}

func NewPizzaUsecase(repo entities.IPizzaRepository) entities.IPizzaUsecase {
	return &pizzaUsecase{
		repo: repo,
	}
}

func (x *pizzaUsecase) GetPizzas() ([]entities.Pizza, error) {
	return x.repo.GetPizzas()
}

func (x *pizzaUsecase) GetPizzaByID(ID string) (entities.Pizza, error) {
	return x.repo.GetPizzaByID(ID)
}

func (x *pizzaUsecase) GetPizzaByName(name string) (entities.Pizza, error) {
	return x.repo.GetPizzaByName(name)
}
