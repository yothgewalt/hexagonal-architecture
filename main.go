package main

import (
	"fmt"
	"hexagonal/handler"
	"hexagonal/logs"
	"hexagonal/repository"
	"hexagonal/service"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	initTimeZone()
	initConfig()
	db := initDatabase()

	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)

	accountRepositoryDB := repository.NewAccountRepositoryDB(db)
	accountService := service.NewAcocuntService(accountRepositoryDB)
	accountHandler := handler.NewAccountHandler(accountService)

	customerService := service.NewCustomerService(customerRepositoryDB)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	router.HandleFunc("/customers/{customerID:[0-9]+}/accounts", accountHandler.GetAccount).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}/accounts", accountHandler.NewAccount).Methods(http.MethodPost)

	logs.Info("Application listening at port: " + viper.GetString("listen.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("listen.port")), router)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDatabase() *sqlx.DB {
	dsn := fmt.Sprintf("%v:%v@/%v",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.database"),
	)

	db, err := sqlx.Open(viper.GetString("database.driver"), dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
