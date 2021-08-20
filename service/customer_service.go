package service

import (
	"database/sql"
	"hexagonal/errs"
	"hexagonal/logs"
	"hexagonal/repository"
)

type customerService struct {
	cRepo repository.CustomerRepository
}

func NewCustomerService(cRepo repository.CustomerRepository) customerService {
	return customerService{cRepo: cRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.cRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	cResponses := []CustomerResponse{}
	for _, customer := range customers {
		cResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		cResponses = append(cResponses, cResponse)
	}

	return cResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.cRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	cResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &cResponse, nil
}
