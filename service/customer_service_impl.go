package service

import (
	"context"
	"database/sql"
	"insectdie/go-restapi-mux/helper"
	"insectdie/go-restapi-mux/model/web"
	"insectdie/go-restapi-mux/repository"

	"github.com/go-playground/validator/v10"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCustomerService(customerRepository repository.CustomerRepository, DB *sql.DB, validate *validator.Validate) CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: customerRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CustomerServiceImpl) FindAll(ctx context.Context) []web.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customers := service.CustomerRepository.FindAll(ctx, tx)

	return helper.ToCustomerResponses(customers)
}
