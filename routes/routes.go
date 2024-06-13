package routes

import (
	"biblioteca/controllers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/create", controllers.Create)
	http.HandleFunc("/delete", controllers.Delete)
}
