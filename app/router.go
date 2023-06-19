package app

import (
	"insectdie/go-restapi-mux/controller"

	"github.com/gorilla/mux"
)

func NewRouter(customerController controller.CustomerController) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/customers", customerController.FindAll).Methods("GET")
	return router
}
