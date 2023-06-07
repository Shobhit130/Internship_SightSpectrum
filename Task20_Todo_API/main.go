package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Todo represents a todo item
type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos map[string]Todo

func main() {
	todos = make(map[string]Todo)

	r := mux.NewRouter()
	r.HandleFunc("/todos", getTodos).Methods("GET")
	r.HandleFunc("/todos", createTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", getTodo).Methods("GET")
	r.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	var todoList []Todo
	for _, todo := range todos {
		todoList = append(todoList, todo)
	}

	jsonResponse(w, http.StatusOK, todoList)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	todos[todo.ID] = todo

	jsonResponse(w, http.StatusCreated, todo)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoID := params["id"]

	todo, exists := todos[todoID]
	if !exists {
		jsonResponse(w, http.StatusNotFound, "Todo not found")
		return
	}

	jsonResponse(w, http.StatusOK, todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoID := params["id"]

	todo, exists := todos[todoID]
	if !exists {
		jsonResponse(w, http.StatusNotFound, "Todo not found")
		return
	}

	var updatedTodo Todo
	err := json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	todo.Title = updatedTodo.Title
	todo.Completed = updatedTodo.Completed
	todos[todoID] = todo

	jsonResponse(w, http.StatusOK, todo)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoID := params["id"]

	_, exists := todos[todoID]
	if !exists {
		jsonResponse(w, http.StatusNotFound, "Todo not found")
		return
	}

	delete(todos, todoID)

	jsonResponse(w, http.StatusOK, "Todo deleted")
}

func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
