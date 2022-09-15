package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
	log.Println("handled a request with success!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{FirstName: "Nikhil", LastName: "Rai"},
		{FirstName: "Ritu", LastName: "Rai"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
