package controller

import "net/http"

type CustomerController interface {
	FindAll(writer http.ResponseWriter, request *http.Request)
}
