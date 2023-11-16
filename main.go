package main

import (
	"encoding/json"
	"log"
	"net/http"
	"todo/api"
	"todo/config"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"request": "success"})
}

func main() {
	// Logging setup
	if err := handleLogging(); err != nil {
		log.Fatalf("Error setting up logging: %v", err)
	}

	// REST API setup
	if err := handleAPI(); err != nil {
		log.Fatalf("Error setting up API: %v", err)
	}
}

func handleLogging() error {
	if err := config.LoadEnvironmentFile(); err != nil {
		return err
	}

	file, err := config.OpenLogFile("./todo.log")
	if err != nil {
		return err
	}

	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	log.Println("Log file created")
	return nil
}

// sets up the REST API endpoints.
func handleAPI() error {
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

	port := config.GetValueFromEnvFile("PORT")
	return http.ListenAndServe(":"+port, nil)
}
