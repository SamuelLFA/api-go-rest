package routes

import (
	"log"
	"net/http"

	"github.com/SamuelLFA/api-go-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/stars", controllers.AllStars)
	r.HandleFunc("/stars/{id}", controllers.FindStarById)
	log.Fatal(http.ListenAndServe(":8000", r))
}
