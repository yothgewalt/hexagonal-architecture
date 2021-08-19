package main

import (
	"fmt"
	"hexagonal/repository"
	"hexagonal/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:!First@14131413@/banking")
	if err != nil {
		panic(err)
	}

	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCustomerService(customerRepository)
	customers, err := customerService.GetCustomers()
	if err != nil {
		panic(err)
	}

	fmt.Println(customers)

}
