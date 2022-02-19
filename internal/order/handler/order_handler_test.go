package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/go-playground/validator"

	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/entities"
	mocks "github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/entities/mocks"
	orderHandler "github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/order/handler"
	"github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/utils"
	internalErrors "github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/utils/errors"
)

func newTestRouter(uc entities.IOrderUsecase) *echo.Echo {
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	orderHandler.NewOrderHandler(e, uc)
	return e
}

var _ = Describe("Http", func() {
	var (
		router   *echo.Echo
		recorder *httptest.ResponseRecorder
		request  *http.Request
		ctrl     *gomock.Controller
		uc       *mocks.MockIOrderUsecase
		order    entities.Order
	)

	orderItem := entities.OrderItem{
		Size:      "Large",
		PizzaName: "Meatfeast",
	}
	order = entities.Order{
		Name:       "jim",
		OrderItems: []entities.OrderItem{orderItem},
	}

	orderID := "d4db3b7c-e6be-4250-bce0-747a3b82949a"

	BeforeEach(func() {
		recorder = httptest.NewRecorder()
		ctrl = gomock.NewController(GinkgoT())
		uc = mocks.NewMockIOrderUsecase(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("POST /api/orders", func() {
		JustBeforeEach(func() {
			router = newTestRouter(uc)
			router.ServeHTTP(recorder, request)
		})

		Context("Happy path", func() {
			BeforeEach(func() {

				uc.EXPECT().CreateOrder(order).Return(orderID, nil)
				bodyBytes, err := json.Marshal(order)
				if err != nil {
					fmt.Println(err)
				}
				request, _ = http.NewRequest("POST", "/api/orders", bytes.NewBuffer(bodyBytes))
				request.Header.Set("Content-Type", "application/json")
			})

			It("returns a 200", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})

			It("returns expected response", func() {
				var responseID interface{}
				if err := json.Unmarshal(recorder.Body.Bytes(), &responseID); err != nil {
					Fail("could not unmarshal response into trip")
				}
				Expect(responseID).To(Equal(orderID))
			})
		})

		When("Validation error parsing body", func() {
			BeforeEach(func() {
				bodyBytes, err := json.Marshal(entities.Order{})
				if err != nil {
					fmt.Println(err)
				}
				request, _ = http.NewRequest("POST", "/api/orders", bytes.NewBuffer(bodyBytes))
				request.Header.Set("Content-Type", "application/json")
			})

			It("returns a 400", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})

			It("returns no response", func() {
				Expect(strings.TrimRight(recorder.Body.String(), "\n")).To(Equal(`{"message":"Key: 'Order.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'Order.OrderItems' Error:Field validation for 'OrderItems' failed on the 'required' tag"}`))
			})
		})

		When("Usecase returns a validation error", func() {
			BeforeEach(func() {
				bodyBytes, err := json.Marshal(order)
				if err != nil {
					fmt.Println(err)
				}
				uc.EXPECT().CreateOrder(order).Return("", &internalErrors.ValidationError{Message: "validation error"})
				request, _ = http.NewRequest("POST", "/api/orders", bytes.NewBuffer(bodyBytes))
				request.Header.Set("Content-Type", "application/json")
			})

			It("returns a 400", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})

			It("returns no response", func() {
				Expect(strings.TrimRight(recorder.Body.String(), "\n")).To(Equal(`{"message":"validation error"}`))
			})
		})

		When("Usecase returns an error", func() {
			BeforeEach(func() {
				bodyBytes, err := json.Marshal(order)
				if err != nil {
					fmt.Println(err)
				}
				uc.EXPECT().CreateOrder(order).Return("", fmt.Errorf("Fail"))
				request, _ = http.NewRequest("POST", "/api/orders", bytes.NewBuffer(bodyBytes))
				request.Header.Set("Content-Type", "application/json")
			})

			It("returns a 500", func() {
				Expect(recorder.Code).To(Equal(http.StatusInternalServerError))
			})

			It("returns expected response", func() {
				Expect(strings.TrimRight(recorder.Body.String(), "\n")).To(Equal(`{"message":"Internal Server Error"}`))
			})
		})
	})
})
