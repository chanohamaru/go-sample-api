package controller

import (
	"fmt"
	"net/http"
)

type TodoController interface {
	GetTodos(w http.ResponseWriter, r *http.Request)
	PostTodo(w http.ResponseWriter, r *http.Request)
}

type todoController struct {
}

func NewTodoController() TodoController {
	return &todoController{}
}

func (tc *todoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "get todos")
}

func (tc *todoController) PostTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applivation/json")
	fmt.Fprintf(w, "post todo")
}
