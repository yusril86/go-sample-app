package routes

import (
	"net/http"

	"github.com/yusril86/go-sample-app/controller"
)

func MapRoutes(server *http.ServeMux) {
	server.HandleFunc("/", controller.HelloWorld())
	server.HandleFunc("/employee", controller.IndexEmployee())

}
