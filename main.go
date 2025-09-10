package main

import (
	"net/http"

	"github.com/yusril86/go-sample-app/database"
	"github.com/yusril86/go-sample-app/routes"
)

func main() {
	database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server)

	http.ListenAndServe(":8080", server)
}
