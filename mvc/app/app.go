package app

import (
	"log"
	"net/http"

	"github.com/novdov/goms/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
