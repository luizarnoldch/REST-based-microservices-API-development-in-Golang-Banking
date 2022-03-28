package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/luizarnoldch/REST-based-microservices-API-development-in-Golang-Banking/domain"
	"github.com/luizarnoldch/REST-based-microservices-API-development-in-Golang-Banking/service"
)

func sanityCheck(){
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment varaible not defined")
	}
}

func Start() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env")
	}
	
	sanityCheck()
	//mux := http.NewServeMux()
	mux := mux.NewRouter()


	//defining database
	dbClient := getDbClient()
	customerRespositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRespositoryDb := domain.NewAccountRepositoryDb(dbClient)

	//define Data

	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb(dbClient))}

	ch := CustomerHandlers{service.NewCustomerService(customerRespositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRespositoryDb)}

	//difine routes

	mux.
		HandleFunc("/customers", ch.getAllCustomers).
		Methods(http.MethodGet)

	mux.
		HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).
		Methods(http.MethodGet)

	mux.
		HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).
		Methods(http.MethodPost)

	mux.
		HandleFunc("/customers/{customer_id:[0-9]+}/account/{customer_id:[0-9]+}", ah.MakeTransaction).
		Methods(http.MethodPost)
	/*
		mux.HandleFunc("/greet", greet).Methods(http.MethodGet)
		mux.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
		//Regex para evaluar solo números

	*/
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	//starting server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s",address,port), mux))
}

func getDbClient() *sqlx.DB {
	driver := "mysql"
	//usuario := "root"
	usuario := os.Getenv("DB_USER")
	//pass := "u1OboD93110614"
	pass := os.Getenv("DB_PASSWD")
	//port := "tcp(localhost:3306)"
	port := os.Getenv("DB_ADDRESS_PORT")
	//table := "banking"
	table := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@%s/%s", usuario, pass, port, table)

	client, err := sqlx.Open(driver, dataSource)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}