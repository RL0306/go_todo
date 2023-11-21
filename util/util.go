package util

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
	"todo/entity"
)

func ExtractIdFromUrl(path string) string {
	pattern := `/todo/(\d+)`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal("Error compiling regex")
	}

	matches := regex.FindStringSubmatch(path)
	if len(matches) < 2 {
		log.Println("Invalid URL format")
	}

	idFromUrl := matches[1]
	if err != nil {
		log.Fatal("Error converting string to int")
	}
	return idFromUrl
}

func WriteTodosToJsonFile(todos []entity.Todo) {
	todosAsJson, _ := json.Marshal(todos)
	err := os.WriteFile("datastore/todos.json", todosAsJson, os.ModePerm)
	if err != nil {
		log.Fatal("Error writing to file")
	}
}
