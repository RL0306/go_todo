package api

import (
	"encoding/json"
	"io"
	"log"
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
	log.Println("Handling GetAllTodos request")
	w.Header().Set("Content-Type", "application/json")
	todosResp, err := json.Marshal(todos)
	if err != nil {
		log.Printf("Error encoding todos to JSON: %s", err.Error())
		responseError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(todosResp)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling CreateTodo request")

	var todoRequest model.TodoRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %s", err.Error())
		responseError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &todoRequest)
	if err != nil {
		log.Printf("Error decoding JSON: %s", err.Error())
		responseError(w, "Bad Request - Invalid JSON", http.StatusBadRequest)
		return
	}

	if todoRequest.Description == "" {
		log.Println("Description cannot be empty")
		responseError(w, "Bad Request - Description cannot be empty", http.StatusBadRequest)
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
		log.Printf("Error encoding JSON: %s", err.Error())
		responseError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(successTodoResponseJSON)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling DeleteTodo request")

	pattern := `/todo/(\d+)`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		log.Printf("Error compiling regex: %s", err.Error())
		responseError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	matches := regex.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		log.Println("Invalid URL format")
		responseError(w, "Bad Request - Invalid URL format", http.StatusBadRequest)
		return
	}

	idFromUrl := matches[1]

	idAsInt, err := strconv.Atoi(idFromUrl)
	if err != nil {
		log.Printf("Invalid ID format: %s", err.Error())
		responseError(w, "Bad Request - Invalid ID format", http.StatusBadRequest)
		return
	}

	if idAsInt < 0 || idAsInt >= len(todos) {
		log.Println("Invalid todo ID")
		responseError(w, "Not Found - Invalid todo ID", http.StatusNotFound)
		return
	}

	todos = append(todos[:idAsInt-1], todos[idAsInt:]...)

	w.Header().Set("Content-Type", "application/json")
	deletedTodoResponse := map[string]string{"message": "Todo successfully deleted"}
	deletedTodoResponseJson, _ := json.Marshal(deletedTodoResponse)
	w.Write(deletedTodoResponseJson)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling UpdateTodo request")

	pattern := `/todo/(\d+)`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		log.Printf("Error compiling regex: %s", err.Error())
		responseError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	idFromUrlMatches := regex.FindStringSubmatch(r.URL.Path)
	if len(idFromUrlMatches) < 2 {
		log.Println("Invalid URL format")
		responseError(w, "Bad Request - Invalid URL format", http.StatusBadRequest)
		return
	}

	idFromUrl := idFromUrlMatches[1]

	idAsInt, err := strconv.Atoi(idFromUrl)
	if err != nil {
		log.Printf("Invalid ID format: %s", err.Error())
		responseError(w, "Bad Request - Invalid ID format", http.StatusBadRequest)
		return
	}

	if idAsInt < 0 || idAsInt >= len(todos) {
		log.Println("Invalid todo ID")
		responseError(w, "Not Found - Invalid todo ID", http.StatusNotFound)
		return
	}

	todo := &todos[idAsInt]
	todo.Completed = true

	w.Header().Set("Content-Type", "application/json")
	updatedTodoResponse := map[string]string{"message": "Todo status updated to completed"}
	updatedTodoResponseJson, err := json.Marshal(updatedTodoResponse)
	if err != nil {
		log.Printf("Error encoding JSON: %s", err.Error())
		responseError(w, "Internal Server Error", http.StatusInternalServerError)
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
