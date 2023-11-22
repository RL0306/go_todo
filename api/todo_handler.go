package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"todo/entity"
)

func TodoHandler(writer http.ResponseWriter, request *http.Request) {
	//do this in main.go
	todoClient := TodoClient{FilePath: "datastore/todos.json"}

	switch request.Method {
	case http.MethodGet:
		GetAllTodoHandler(todoClient)
	}

}

type TodoClientInterface interface {
	GetAllTodos() []entity.Todo
}

func GetAllTodoHandler(todoClient TodoClientInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling GetAllTodos request")
		todos := todoClient.GetAllTodos()
		fmt.Println(todos)
		todosResponse, _ := json.Marshal(todos)
		w.Header().Set("Content-Type", "application/json")
		w.Write(todosResponse)

	}
}

type TodoClient struct {
	FilePath string
}

func (c TodoClient) GetAllTodos() []entity.Todo {
	var Todos []entity.Todo

	file, err := os.Open(c.FilePath)
	if err != nil {
		log.Fatal("Error opening file")
	}

	defer file.Close()

	byteResult, err := io.ReadAll(file)

	if err != nil {
		log.Fatal("Error reading file")
		return nil
	}

	err = json.Unmarshal(byteResult, &Todos)
	if err != nil {
		return nil
	}

	return Todos
}

//func getAllTodos(writer http.ResponseWriter, request *http.Request) {
//	log.Println("Handling GetAllTodos request")
//
//	//reading a file then parsing the json and returning back a slice
//	todos := service.GetAllTodos()
//	todosResponse, err := json.Marshal(todos)
//
//	if err != nil {
//		log.Fatal("Error marshalling todos data")
//	}
//
//	writer.Header().Set("Content-Type", "application/json")
//	writer.Write(todosResponse)
//
//}
//
//func createTodo(writer http.ResponseWriter, request *http.Request) {
//	log.Println("Handling CreateTodo request")
//
//	status := service.CreateTodo(request.Body)
//	writer.Header().Set("Content-Type", "application/json")
//
//	if status {
//		success := map[string]string{"status": "Successfully created todo"}
//		successJson, _ := json.Marshal(success)
//		writer.Write(successJson)
//	} else {
//		failure := map[string]string{"status": "Failed creating todo"}
//		failureJson, _ := json.Marshal(failure)
//		writer.Write(failureJson)
//	}
//
//}
//
//func deleteTodo(writer http.ResponseWriter, request *http.Request) {
//	log.Println("Handling DeleteTodo request")
//
//	id := util.ExtractIdFromUrl(request.URL.Path)
//
//	status := service.DeleteTodo(id)
//
//	writer.Header().Set("Content-Type", "application/json")
//	if status {
//		success := map[string]string{"status": "Successfully deleted todo"}
//		successJson, _ := json.Marshal(success)
//		writer.Write(successJson)
//	} else {
//		failed := map[string]string{"status": "Could not delete todo"}
//		failedJson, _ := json.Marshal(failed)
//		writer.Write(failedJson)
//	}
//}
//
//func updateTodo(writer http.ResponseWriter, request *http.Request) {
//	log.Println("Handling UpdateTodo request")
//	id := util.ExtractIdFromUrl(request.URL.Path)
//
//	status := service.UpdateTodo(id)
//	writer.Header().Set("Content-Type", "application/json")
//
//	if status {
//		success := map[string]string{"status": "Successfully updated todo"}
//		successJson, _ := json.Marshal(success)
//		writer.Write(successJson)
//	} else {
//		failed := map[string]string{"status": "Could not update todo"}
//		failedJson, _ := json.Marshal(failed)
//		writer.Write(failedJson)
//	}
//}
//
//func getTodoById(writer http.ResponseWriter, request *http.Request) {
//	log.Println("Handling GetTodoById request")
//	id := util.ExtractIdFromUrl(request.URL.Path)
//
//	todo, err := service.GetTodoById(id)
//	writer.Header().Set("Content-Type", "application/json")
//	if err != nil {
//		errResponse, _ := json.Marshal(err)
//		writer.Write(errResponse)
//	} else {
//		todoResponse, _ := json.Marshal(todo)
//		writer.Write(todoResponse)
//	}
//}
