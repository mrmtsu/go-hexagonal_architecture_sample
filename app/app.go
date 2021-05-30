package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrmtsu/go-api-rest/domain"
	"github.com/mrmtsu/go-api-rest/service"
)

func Start() {

	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe(":8000", router))
}
