package main

import (
	"net/http"

	"github.com/chanohamaru/go-sample-api/controller"
)

var tc = controller.NewTodoController()
var ro = controller.NewRouter(tc)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/todos/", ro.HandleTodosRequest)
	server.ListenAndServe()
}
