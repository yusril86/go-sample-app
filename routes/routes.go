package routes

import (
	"database/sql"
	"net/http"

	"github.com/yusril86/go-sample-app/controller"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.HelloWorld())
	server.HandleFunc("/employee", controller.IndexEmployee(db))
	server.HandleFunc("/employee/create", controller.CreateEmployee(db))

}
