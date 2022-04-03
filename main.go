package main

import (
	"api-rest-simple/models"
	"api-rest-simple/routes"
	"fmt"
)

func main() {
	fmt.Println("Iniciando servidor porta 8080")

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Gabriel", Historia: "História de Gabriel"},
		{Id: 2, Nome: "Juliana", Historia: "História de Juliana"},
	}

	routes.HandleRequest()
}
