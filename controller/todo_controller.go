package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/chanohamaru/go-sample-api/controller/dto"
	"github.com/chanohamaru/go-sample-api/model/entity"
	"github.com/chanohamaru/go-sample-api/model/repository"
)

type TodoController interface {
	GetTodos(w http.ResponseWriter, r *http.Request)
	PostTodo(w http.ResponseWriter, r *http.Request)
}

type todoController struct {
	tr repository.TodoRepository
}

func NewTodoController(tr repository.TodoRepository) TodoController {
	return &todoController{tr}
}

func (tc *todoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := tc.tr.GetTodos()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var todoResponses []dto.TodoResponse
	for _, v := range todos {
		todoResponses = append(todoResponses, dto.TodoResponse{Id: v.Id, Title: v.Title, Content: v.Content})
	}

	var todosResponse dto.TodosResponse
	todosResponse.Todos = todoResponses

	output, _ := json.MarshalIndent(todosResponse.Todos, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (tc *todoController) PostTodo(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var todoRequest dto.TodoRequest
	json.Unmarshal(body, &todoRequest)

	todo := entity.TodoEntity{Title: todoRequest.Title, Content: todoRequest.Content}
	id, err := tc.tr.InsertTodo(todo)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
	w.WriteHeader(201)
	fmt.Fprintf(w, r.Host+r.URL.Path+strconv.Itoa(id))
}
