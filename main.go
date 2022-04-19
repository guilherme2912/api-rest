package main

import (
	"fmt"

	"github.com/guilherme2912/go-rest-api/models"
	"github.com/guilherme2912/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Executando servidor")
	routes.HandleResquet()
}
