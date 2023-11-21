package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"todo/config"
	"todo/entity"
	"todo/util"
)

func GetAllTodos() []entity.Todo {
	var Todos []entity.Todo

	file, err := os.Open("datastore/todos.json")
	if err != nil {
		log.Fatal("Error opening file")
	}

	defer file.Close()

	byteResult, err := io.ReadAll(file)

	if err != nil {
		log.Fatal("Error reading file")
	}

	json.Unmarshal(byteResult, &Todos)

	return Todos

}

func CreateTodo(bodyRequest io.ReadCloser) bool {
	bodyBytes, err := io.ReadAll(bodyRequest)
	if err != nil {
		log.Println("Error reading in request")
		return false
	}

	var todoRequest entity.TodoRequest

	err = json.Unmarshal(bodyBytes, &todoRequest)

	if err != nil {
		log.Println("Error parsing body to todo")
		return false
	}

	todos := GetAllTodos()
	id, _ := strconv.Atoi(config.GetValueFromEnvFile("ID_VALUE"))

	todo := entity.Todo{
		Id:          id,
		Description: todoRequest.Description,
		Completed:   false,
	}

	os.Setenv("ID_VALUE", strconv.Itoa(id+1))

	todos = append(todos, todo)
	util.WriteTodosToJsonFile(todos)

	return true
}

func GetTodoById(id string) (entity.Todo, map[string]string) {
	file, err := os.Open("datastore/todos.json")
	if err != nil {
		log.Println("Error opening file")
	}

	defer file.Close()

	byteResult, err := io.ReadAll(file)
	var todos []entity.Todo

	err = json.Unmarshal(byteResult, &todos)
	if err != nil {
		log.Println("Error unmarshalling json data")
	}

	idAsInt, err := strconv.Atoi(id)

	if err != nil {
		log.Println("Error converting value to int")
	}

	for _, todo := range todos {
		if todo.Id == idAsInt {
			return todo, nil
		}
	}

	return entity.Todo{}, map[string]string{"status": "Could not find todo"}

}

func DeleteTodo(id string) bool {
	idAsInt, err := strconv.Atoi(id)

	if err != nil {
		log.Printf("Error converting %d to int", idAsInt)
	}

	todos := GetAllTodos()

	if idAsInt > len(todos) || idAsInt < 0 {
		log.Printf("Todo with ID %d does not exist", idAsInt)
		return false
	}

	for i, todo := range todos {
		if idAsInt == todo.Id {
			todos = append(todos[:i], todos[i+1:]...)
		}
	}
	util.WriteTodosToJsonFile(todos)
	fmt.Println(todos)

	return true
}

func UpdateTodo(id string) bool {
	idAsInt, err := strconv.Atoi(id)

	if err != nil {
		log.Printf("Error converting %d to int", idAsInt)
	}

	todos := GetAllTodos()

	if idAsInt > len(todos) || idAsInt < 0 {
		log.Printf("Todo with ID %d does not exist", idAsInt)
		return false
	}

	for _, todo := range todos {
		if todo.Id == idAsInt {
			var todoToUpdate = todo
			todoToUpdate.Completed = true
			util.WriteTodosToJsonFile(todos)
			return true
		}
	}
	return false
}
