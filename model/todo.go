package model

type Todo struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// TodoRequest This is for when a user creates a new Todo_ as they won't be setting id and completed
type TodoRequest struct {
	Description string `json:"description"`
}
