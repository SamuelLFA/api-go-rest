package main

import (
	"fmt"

	"github.com/SamuelLFA/api-go-rest/models"
	"github.com/SamuelLFA/api-go-rest/routes"
)

func main() {
	models.Stars = []models.Star{
		{Name: "Nome 1", Story: "Story"},
		{Name: "Nome 2", Story: "Story2"},
	}

	fmt.Println("Iniciando o servidor Rest em Go")
	routes.HandleRequest()
}
