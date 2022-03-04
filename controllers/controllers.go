package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SamuelLFA/api-go-rest/database"
	"github.com/SamuelLFA/api-go-rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllStars(w http.ResponseWriter, r *http.Request) {
	var stars []models.Star
	database.DB.Find(&stars)
	json.NewEncoder(w).Encode(stars)
}

func FindStarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var star models.Star

	database.DB.First(&star, id)
	json.NewEncoder(w).Encode(star)
}
