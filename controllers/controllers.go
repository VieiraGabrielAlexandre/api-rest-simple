package controllers

import (
	"api-rest-simple/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}

func CapturarPersonalidade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	for _, item := range models.Personalidades {
		if strconv.Itoa(item.Id) == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Personalidade{})
}
