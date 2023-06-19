package controller

import (
	"insectdie/go-restapi-mux/helper"
	"insectdie/go-restapi-mux/model/web"
	"insectdie/go-restapi-mux/service"
	"net/http"
)

type CustomerControllerImpl struct {
	CustomerService service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) CustomerController {
	return &CustomerControllerImpl{
		CustomerService: customerService,
	}
}

func (controller *CustomerControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	customerResponses := controller.CustomerService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
