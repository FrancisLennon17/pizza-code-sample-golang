package order_test

import (
	"errors"

	gomock "github.com/golang/mock/gomock"
	"github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/entities"
	"github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/entities/mocks"
	orderUsecase "github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/order"
)

var errMock = errors.New("error")

var _ = Describe("order usecase", func() {
	var (
		ctrl           *gomock.Controller
		order          entities.Order
		orderRepo      *mocks.MockIOrderRepository
		pizzaUC        *mocks.MockIPizzaUsecase
		orderUC        entities.IOrderUsecase
		pizza          entities.Pizza
		orderID        string
		orderCreatedID string
		err            error
	)

	orderItem := entities.OrderItem{
		Size:      "Large",
		PizzaName: "Meatfeast",
	}
	order = entities.Order{
		Name:       "jim",
		OrderItems: []entities.OrderItem{orderItem},
	}

	pizza = entities.Pizza{
		ID:          "",
		Name:        "Meatfeast",
		Price:       12.50,
		Ingredients: pq.StringArray{"Cheese", "Meat", "Sauce"},
	}

	orderID = "d4db3b7c-e6be-4250-bce0-747a3b82949a"

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		orderRepo = mocks.NewMockIOrderRepository(ctrl)
		pizzaUC = mocks.NewMockIPizzaUsecase(ctrl)
		orderUC = orderUsecase.NewOrderUsecase(orderRepo, pizzaUC)
	})

	Describe("Create Order", func() {
		When("Happy path", func() {
			BeforeEach(func() {
				pizzaUC.EXPECT().GetPizzaByName(orderItem.PizzaName).Return(pizza, nil)
				orderRepo.EXPECT().CreateOrder(order).Return(orderID, nil)
				orderCreatedID, err = orderUC.CreateOrder(order)
			})

			It("Does not return error", func() {
				Expect(err).To(BeNil())
			})

			It("returns expected order ID", func() {
				Expect(orderCreatedID).To(Equal(orderID))
			})
		})

		When("pizza usecase returns error", func() {
			BeforeEach(func() {
				pizzaUC.EXPECT().GetPizzaByName(orderItem.PizzaName).Return(pizza, errMock)
				orderCreatedID, err = orderUC.CreateOrder(order)
			})

			It("Does return error", func() {
				Expect(err).To(Equal(errMock))
			})
		})

		When("order repo returns error", func() {
			BeforeEach(func() {
				pizzaUC.EXPECT().GetPizzaByName(orderItem.PizzaName).Return(pizza, nil)
				orderRepo.EXPECT().CreateOrder(order).Return(orderID, errMock)
				orderCreatedID, err = orderUC.CreateOrder(order)
			})

			It("Does return error", func() {
				Expect(err).To(Equal(errMock))
			})
		})
	})
})
