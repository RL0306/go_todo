package main

import (
	"encoding/json"
	"net/http"
	"todo/api"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"request": "success"})
}

func main() {
	http.HandleFunc("/test", TestHandler)
	http.HandleFunc("/todo", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			api.GetAllTodos(writer, request)
		case http.MethodPost:
			api.CreateTodo(writer, request)
		}
	})

	http.HandleFunc("/todo/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodDelete:
			api.DeleteTodo(writer, request)
		case http.MethodPatch:
			api.UpdateTodo(writer, request)
		}
	})

	http.ListenAndServe(":8080", nil)
}
