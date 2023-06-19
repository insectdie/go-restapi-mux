package helper

import (
	"insectdie/go-restapi-mux/model/domain"
	"insectdie/go-restapi-mux/model/web"
)

func ToCustomerResponse(category domain.Customer) web.CustomerResponse {
	return web.CustomerResponse{
		Id:          category.Id,
		Name:        category.Name,
		Dob:         category.Dob,
		Phone:       category.Phone,
		Nationality: category.Nationality,
		Email:       category.Email,
		Family:      category.Family,
	}
}

func ToCustomerResponses(categories []domain.Customer) []web.CustomerResponse {
	var customerResponses []web.CustomerResponse
	for _, customer := range categories {
		customerResponses = append(customerResponses, ToCustomerResponse(customer))
	}
	return customerResponses
}
