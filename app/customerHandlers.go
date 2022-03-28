package app

import (
	"encoding/json"
	//"encoding/xml"
	//"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luizarnoldch/REST-based-microservices-API-development-in-Golang-Banking/service"
)

/*
type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}
*/

type CustomerHandlers struct {
	service service.CustomerService
}

/*
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")
}
*/

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	/*
		customers := []Customer{
			{"Ashish", "New Delhi", "11075"},
			{"Ashish", "New Delhi", "11075"},
		}
	*/
	
	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomer(status)
	
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w,http.StatusOK, customers)
	}
	/*
		if r.Header.Get("Content-Type") == "application/xml" {
			w.Header().Add("Content-Type", "application/xml")
			xml.NewEncoder(w).Encode(customers)
		} else {
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(customers)
		}
	*/
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		/*
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(err.Code)
			//fmt.Fprintf(w, err.Message)
			json.NewEncoder(w).Encode(err.AsMessage())
		*/
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		/*
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customer)
		*/
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

/*
func createCustomer(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Post request received")
}
*/
