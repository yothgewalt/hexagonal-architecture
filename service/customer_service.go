package service

import (
	"hexagonal/repository"
	"log"
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
		log.Println(err)
		return nil, err
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
	return nil, nil
}
