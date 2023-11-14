package api

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"todo/model"
)

var todos = []model.Todo{
	{Id: 1, Description: "This is my first todo", Completed: false},
	{Id: 2, Description: "This is my second todo", Completed: false},
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todosResp, _ := json.Marshal(todos)
	w.Write(todosResp)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todoRequest model.TodoRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responseError(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &todoRequest)
	if err != nil {
		responseError(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	if todoRequest.Description == "" {
		responseError(w, "Description cannot be empty", http.StatusBadRequest)
		return
	}

	newTodo := model.Todo{
		Id:          len(todos) + 1,
		Description: todoRequest.Description,
		Completed:   false,
	}

	todos = append(todos, newTodo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	successTodoResponse := map[string]string{"message": "Todo successfully created"}
	successTodoResponseJSON, err := json.Marshal(successTodoResponse)
	if err != nil {
		responseError(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.Write(successTodoResponseJSON)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	pattern := `/todo/(\d+)`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		responseError(w, "Error compiling regex", http.StatusInternalServerError)
		return
	}

	matches := regex.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		responseError(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	idFromUrl := matches[1]

	idAsInt, err := strconv.Atoi(idFromUrl)
	if err != nil {
		responseError(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	if idAsInt < 0 || idAsInt >= len(todos) {
		responseError(w, "Invalid todo ID", http.StatusNotFound)
		return
	}

	todos = append(todos[:idAsInt-1], todos[idAsInt:]...)

	w.Header().Set("Content-Type", "application/json")
	deletedTodoResponse := map[string]string{"message": "Todo successfully deleted"}
	deletedTodoResponseJson, _ := json.Marshal(deletedTodoResponse)
	w.Write(deletedTodoResponseJson)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	pattern := `/todo/(\d+)`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		responseError(w, "Error compiling regex", http.StatusInternalServerError)
		return
	}

	idFromUrlMatches := regex.FindStringSubmatch(r.URL.Path)
	if len(idFromUrlMatches) < 2 {
		responseError(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	idFromUrl := idFromUrlMatches[1]

	idAsInt, err := strconv.Atoi(idFromUrl)
	if err != nil {
		responseError(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	if idAsInt < 0 || idAsInt >= len(todos) {
		responseError(w, "Invalid todo ID", http.StatusNotFound)
		return
	}

	todo := &todos[idAsInt]
	todo.Completed = true

	w.Header().Set("Content-Type", "application/json")
	updatedTodoResponse := map[string]string{"message": "Todo status updated to completed"}
	updatedTodoResponseJson, err := json.Marshal(updatedTodoResponse)
	if err != nil {
		responseError(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.Write(updatedTodoResponseJson)
}

func responseError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorResponse := map[string]string{"error": message}
	errorResponseJSON, _ := json.Marshal(errorResponse)
	w.Write(errorResponseJSON)
}
