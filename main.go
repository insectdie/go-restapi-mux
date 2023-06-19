package main

import (
	"insectdie/go-restapi-mux/app"
	"insectdie/go-restapi-mux/controller"
	"insectdie/go-restapi-mux/helper"
	"insectdie/go-restapi-mux/repository"
	"insectdie/go-restapi-mux/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository, db, validate)
	customerController := controller.NewCustomerController(customerService)
	router := app.NewRouter(customerController)

	err := http.ListenAndServe(":3000", router)
	helper.PanicIfError(err)
}
