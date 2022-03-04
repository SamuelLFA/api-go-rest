package main

import (
	"fmt"

	"github.com/SamuelLFA/api-go-rest/database"
	"github.com/SamuelLFA/api-go-rest/models"
	"github.com/SamuelLFA/api-go-rest/routes"
)

func main() {
	models.Stars = []models.Star{
		{Id: 1, Name: "Nome 1", Story: "Story"},
		{Id: 2, Name: "Nome 2", Story: "Story2"},
	}
	database.ConectWithDatabase()
	fmt.Println("Iniciando o servidor Rest em Go")
	routes.HandleRequest()
}
