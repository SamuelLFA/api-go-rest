package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SamuelLFA/api-go-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllStars(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Stars)
}
