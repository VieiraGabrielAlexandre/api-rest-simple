package routes

import (
	"api-rest-simple/controllers"
	"api-rest-simple/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.CapturarPersonalidade).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.CriarPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletarPersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditarPersonalidade).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
