package main

import (
	"net/http"

	"github.com/yusril86/go-sample-app/database"
	"github.com/yusril86/go-sample-app/routes"
)

func main() {
	db := database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server, db)

	http.ListenAndServe(":8080", server)
}
