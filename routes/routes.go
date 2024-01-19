package routes

import (
	"net/http"

	"github.com/farismasud/crud-employee-go/controller"
)

func MapRoutes(server *http.ServeMux) {
	server.HandleFunc("/", controller.NewHelloWorldController())
}