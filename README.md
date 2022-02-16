# pizza-code-sample-golang
Loosely following the challenge listed here https://github.com/AmbulnzLLC/fullstack-challenge/tree/master/backend

https://www.loom.com/share/23a2062bdb6e4f9794eeff04cb3ff8b0

## Included

### API Endpoints
- GET /api/pizzas returns a list of pizzas 
- POST /api/orders creates an order with a given pizza
- GET /api/orders/:id fetches the details of a given order

## Introduction

## Getting started
### Local Development
- Ensure docker is installed and running
- Run `docker-compose up` from the root of this project
- Ensure golang is installed
- Run `go run cmd/main.go` from the root of this project

### Database migrations

I've used Gorm's AutoMigrate function to create the tables based on the entity structs. For production ready applications, database migrations scripts/tools should be used.

### Seeding the database

Insert pizzas into the database for testing

```
INSERT INTO pizza(
	name, price, ingredients)
	VALUES ('Margherita', 13, ARRAY['tomato', 'mozzarella']);
```

### Swagger UI
When running the service, the openapi3 spec is hosted locally with swagger ui where you can view the endpoint's arguments and possible responses, as well as test the endpoints without the need to use Curl/Postman/Insomnia etc. This was done using swaggo/echo-swagger middleware https://github.com/swaggo/echo-swagger

`http://localhost:8080/swagger/index.html`

### To create/update the swagger file
`go install github.com/swaggo/swag/cmd/swag@latest`
`swag init -g cmd/main.go`

## Architecture
This repository has been based on the Clean Architecture (Robert Martin) template found here https://github.com/bxcodec/go-clean-arch

## ORM
I've used https://gorm.io/ as the ORM for this project.

## Testing
A sample of unit tests have been added at each layer of the service (Handler, Usecase, Repository)
Run unit tests with `go test ./...`

BDD style testing with
- Ginkgo (Testing Framework): https://onsi.github.io/ginkgo/
- Gomega (Matcher Library): https://github.com/onsi/gomega


- GoMock (Mocking): https://github.com/golang/mock
Interfaces are mocked automatically via the comment above the interface that looks like the following
`go:generate mockgen -destination=./mocks/mock_order_repository.go -package=mocks github.com/francislennon17/pizza-code-sample-golang/m/v2/internal/entities IOrderRepository`

## Linting
Run `golangci-lint run` to check for linting errors

## Logging
Using the built in logging functionality in golang & the logging middleware that the echo framework provides

## Notes
When fetching or creating orders, I've purposely seperated the responsibilities so that the order usecase has to reach out to the pizza usecase. This could have been done through the database instead but I wanted to show how each usecase would communicate in more complex scenarios. 

## TO DO
- Better unit test coverage
- DB migraitions
- Entity DTOs to better handle hiding of internal attributes from API consumers
- API testing
- Integration testing
- Better config management
- Swagger Doc validation
