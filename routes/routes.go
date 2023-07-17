package routes

import (
	"log"
	"net/http"

	"github.com/marcelogbrito/go-rest-api/controllers"
)

func HandleResquest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
