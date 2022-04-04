package controllers

import (
	"api-rest-simple/database"
	"api-rest-simple/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade
	database.DB.Find(&personalidades)
	json.NewEncoder(w).Encode(personalidades)
}

func CapturarPersonalidade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func CriarPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade
	_ = json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Create(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}

func DeletarPersonalidade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	database.DB.Delete(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}

func EditarPersonalidade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)

	_ = json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}
