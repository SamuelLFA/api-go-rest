package routes

import (
	"log"
	"net/http"

	"github.com/SamuelLFA/api-go-rest/controllers"
	"github.com/SamuelLFA/api-go-rest/middleware"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/stars", controllers.AllStars).Methods("Get")
	r.HandleFunc("/stars", controllers.CreateNewStar).Methods("Post")
	r.HandleFunc("/stars/{id}", controllers.FindStarById).Methods("Get")
	r.HandleFunc("/stars/{id}", controllers.DeleteStar).Methods("Delete")
	r.HandleFunc("/stars/{id}", controllers.UpdateStar).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
