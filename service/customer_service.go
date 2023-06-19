package service

import (
	"context"
	"insectdie/go-restapi-mux/model/web"
)

type CustomerService interface {
	FindAll(ctx context.Context) []web.CustomerResponse
}
