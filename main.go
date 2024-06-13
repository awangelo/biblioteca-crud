package main

import (
	"biblioteca/database"
	"biblioteca/routes"
	"log"
	"net/http"
)

func main() {
	err := database.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Erro ao conectar com a database: %v", err)
	}

	routes.SetupRoutes()

	http.ListenAndServe(":8000", nil)
}
