package sql_test

import (
	"errors"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"

	mockDB "github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/db/mock"
	"github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/entities"
	pizzaSQL "github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/pizza/repository/sql"
	internalErrors "github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/utils/errors"
)

var _ = Describe("Pizza Repository", func() {
	pizzaID := "f375e3f8-a09c-49f6-be62-b77c5e1e94e9"
	pizzaName := "Meatfeast"
	errMock := errors.New("error")
	validationErrName := &internalErrors.ValidationError{Message: "Could not find pizza with provided name"}
	validationErrID := &internalErrors.ValidationError{Message: "Could not find pizza with provided ID"}
	pizza := entities.Pizza{
		ID:          pizzaID,
		Name:        pizzaName,
		Price:       12.50,
		Ingredients: pq.StringArray{"Cheese", "Meat", "Sauce"},
	}

	db, mock := mockDB.NewMockDB()
	pizzaRepo := pizzaSQL.NewPizzaRepository(db)
	Describe("Get Pizzas", func() {
		var (
			pizzas []entities.Pizza
			err    error
		)
		Context("Success", func() {

			BeforeEach(func() {
				mock.
					ExpectQuery("SELECT.*").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "ingredients"}).
						AddRow(pizzaID, "Meatfeast", 12.50, pq.StringArray{"Cheese", "Meat", "Sauce"}))
				pizzas, err = pizzaRepo.GetPizzas()
			})
			It("Does not return an error", func() {
				Expect(err).To(BeNil())
			})
			It("Returns expected response", func() {
				Expect(pizzas).Should(ConsistOf([]entities.Pizza{pizza}))
			})
		})

		Context("Failure", func() {
			BeforeEach(func() {
				mock.
					ExpectQuery("SELECT.*").
					WillReturnError(errMock)
				pizzas, err = pizzaRepo.GetPizzas()
			})
			It("Does return an error", func() {
				Expect(err).To(Equal(errMock))
			})
		})
	})
	Describe("Get Pizza by ID", func() {
		var (
			pizza entities.Pizza
			err   error
		)
		When("Success", func() {
			BeforeEach(func() {
				mock.
					ExpectQuery("SELECT.*").
					WithArgs(pizzaID).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "ingredients"}).
						AddRow(pizzaID, "Meatfeast", 12.50, pq.StringArray{"Cheese", "Meat", "Sauce"}))
				pizza, err = pizzaRepo.GetPizzaByID(pizzaID)
			})
			It("Does not return an error", func() {
				Expect(err).To(BeNil())
			})
			It("Returns expected response", func() {
				Expect(pizza).To(Equal(pizza))
			})
		})

		When("No Rows", func() {
			BeforeEach(func() {
				mock.
					ExpectQuery("SELECT.*").
					WillReturnError(gorm.ErrRecordNotFound)
				pizza, err = pizzaRepo.GetPizzaByID(pizzaID)
			})
			It("Does return an error", func() {
				Expect(err).To(Equal(validationErrID))
			})
		})

		When("Failure", func() {
			BeforeEach(func() {
				mock.
					ExpectQuery("SELECT.*").
					WillReturnError(errMock)
				pizza, err = pizzaRepo.GetPizzaByID(pizzaID)
			})
			It("Does return an error", func() {
				Expect(err).To(Equal(errMock))
			})
		})
	})

	Describe("Get Pizza by Name", func() {
		var (
			pizza entities.Pizza
			err   error
		)
		When("Success", func() {
			BeforeEach(func() {
				mock.
					ExpectQuery("SELECT.*").
					WithArgs(pizzaName).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "ingredients"}).
						AddRow(pizzaID, "Meatfeast", 12.50, pq.StringArray{"Cheese", "Meat", "Sauce"}))
				pizza, err = pizzaRepo.GetPizzaByName(pizzaName)
			})
			It("Does not return an error", func() {
				Expect(err).To(BeNil())
			})
			It("Returns expected response", func() {
				Expect(pizza).To(Equal(pizza))
			})
		})

		When("No Rows", func() {
			BeforeEach(func() {
				mock.
					ExpectQuery("SELECT.*").
					WillReturnError(gorm.ErrRecordNotFound)
				pizza, err = pizzaRepo.GetPizzaByName(pizzaName)
			})
			It("Does return an error", func() {
				Expect(err).To(Equal(validationErrName))
			})
		})

		When("Failure", func() {
			BeforeEach(func() {
				mock.
					ExpectQuery("SELECT.*").
					WillReturnError(errMock)
				pizza, err = pizzaRepo.GetPizzaByName(pizzaName)
			})
			It("Does return an error", func() {
				Expect(err).To(Equal(errMock))
			})
		})
	})
})
