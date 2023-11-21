package main

import (
	"net/http"
	"todo/api"
	"todo/config"
)

func main() {

	//this is for get all and creating todos
	http.HandleFunc("/todo", api.TodoHandler)

	//this below is for get by id/update/delete
	//in these requests we will be passing in an id after the '/todo/ and that is why we need this
	http.HandleFunc("/todo/", api.TodoHandler)
	config.InitialiseLoggingConfig("./todo.log")

	port := config.GetValueFromEnvFile("PORT")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return
	}
}
