package app

import (
	"log"
	"net/http"
)

func Start() {
	//defining routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
