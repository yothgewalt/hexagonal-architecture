package handler

import (
	"encoding/json"
	"hexagonal/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customerHandler struct {
	cService service.CustomerService
}

func NewCustomerHandler(cService service.CustomerService) customerHandler {
	return customerHandler{cService: cService}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.cService.GetCustomers()
	if err != nil {
		handlerError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	customer, err := h.cService.GetCustomer(customerID)
	if err != nil {
		handlerError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
