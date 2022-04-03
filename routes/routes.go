package routes

import (
	"api-rest-simple/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/personalidades/{id}", controllers.CapturarPersonalidade).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
