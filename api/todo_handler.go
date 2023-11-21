package api

import (
	"encoding/json"
	"log"
	"net/http"
	"todo/service"
	"todo/util"
)

func TodoHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/todo" {
		switch request.Method {
		case http.MethodGet:
			getAllTodos(writer, request)
		case http.MethodPost:
			createTodo(writer, request)
		}
	} else {
		switch request.Method {
		case http.MethodDelete:
			deleteTodo(writer, request)
		case http.MethodPatch:
			updateTodo(writer, request)
		case http.MethodGet:
			getTodoById(writer, request)
		}
	}
}

func getAllTodos(writer http.ResponseWriter, request *http.Request) {
	log.Println("Handling GetAllTodos request")

	//reading a file then parsing the json and returning back a slice
	todos := service.GetAllTodos()
	todosResponse, err := json.Marshal(todos)

	if err != nil {
		log.Fatal("Error marshalling todos data")
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(todosResponse)

}

func createTodo(writer http.ResponseWriter, request *http.Request) {
	log.Println("Handling CreateTodo request")

	status := service.CreateTodo(request.Body)
	writer.Header().Set("Content-Type", "application/json")

	if status {
		success := map[string]string{"status": "Successfully created todo"}
		successJson, _ := json.Marshal(success)
		writer.Write(successJson)
	} else {
		failure := map[string]string{"status": "Failed creating todo"}
		failureJson, _ := json.Marshal(failure)
		writer.Write(failureJson)
	}

}

func deleteTodo(writer http.ResponseWriter, request *http.Request) {
	log.Println("Handling DeleteTodo request")

	id := util.ExtractIdFromUrl(request.URL.Path)

	status := service.DeleteTodo(id)

	writer.Header().Set("Content-Type", "application/json")
	if status {
		success := map[string]string{"status": "Successfully deleted todo"}
		successJson, _ := json.Marshal(success)
		writer.Write(successJson)
	} else {
		failed := map[string]string{"status": "Could not delete todo"}
		failedJson, _ := json.Marshal(failed)
		writer.Write(failedJson)
	}
}

func updateTodo(writer http.ResponseWriter, request *http.Request) {
	log.Println("Handling UpdateTodo request")
	id := util.ExtractIdFromUrl(request.URL.Path)

	status := service.UpdateTodo(id)
	writer.Header().Set("Content-Type", "application/json")

	if status {
		success := map[string]string{"status": "Successfully updated todo"}
		successJson, _ := json.Marshal(success)
		writer.Write(successJson)
	} else {
		failed := map[string]string{"status": "Could not update todo"}
		failedJson, _ := json.Marshal(failed)
		writer.Write(failedJson)
	}
}

func getTodoById(writer http.ResponseWriter, request *http.Request) {
	log.Println("Handling GetTodoById request")
	id := util.ExtractIdFromUrl(request.URL.Path)

	todo, err := service.GetTodoById(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		errResponse, _ := json.Marshal(err)
		writer.Write(errResponse)
	} else {
		todoResponse, _ := json.Marshal(todo)
		writer.Write(todoResponse)
	}
}
