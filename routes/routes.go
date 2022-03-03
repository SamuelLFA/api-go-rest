package routes

import (
	"log"
	"net/http"

	"github.com/SamuelLFA/api-go-rest/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/stars", controllers.AllStars)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
