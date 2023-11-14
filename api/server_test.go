package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/model"
)

func TestGetAllTodos(t *testing.T) {

	test := struct {
		name     string
		expected string
		code     int
	}{
		name:     "Test get all todos endpoint",
		expected: `[{"id":1,"description":"This is my first todo","completed":false},{"id":2,"description":"This is my second todo","completed":false}]`,
		code:     200,
	}

	request := httptest.NewRequest(http.MethodGet, "/todos", nil)
	response := httptest.NewRecorder()

	GetAllTodos(response, request)

	responseBody := response.Body.String()
	statusCode := response.Result().StatusCode

	if responseBody != test.expected {
		t.Errorf("got %v, want %v", responseBody, test.expected)
	}

	if statusCode != test.code {
		t.Errorf("got %v, want %v", response.Result().StatusCode, test.code)
	}
}

func TestCreateTodo(t *testing.T) {
	tests := []struct {
		name     string
		request  model.TodoRequest
		expected string
		code     int
	}{
		{
			name:     "Testing create todo with successful request",
			request:  model.TodoRequest{Description: "This is a test todo"},
			expected: `{"message":"Todo successfully created"}`,
			code:     201,
		},
		{
			name:     "Testing create todo with empty description",
			request:  model.TodoRequest{Description: ""},
			expected: `{"error":"Description cannot be empty"}`,
			code:     400,
		},
	}

	for _, test := range tests {
		todoJsonRequest, _ := json.Marshal(test.request)
		request := httptest.NewRequest(http.MethodPost, "/todo", bytes.NewBuffer(todoJsonRequest))

		response := httptest.NewRecorder()

		CreateTodo(response, request)

		responseBody := response.Body.String()
		statusCode := response.Result().StatusCode

		if responseBody != test.expected {
			t.Errorf("got %v, want %v", response.Body.String(), test.expected)
		}

		if statusCode != test.code {
			t.Errorf("got %v, want %v", response.Result().StatusCode, test.code)
		}
	}
}

func TestDeleteTodo(t *testing.T) {
	tests := []struct {
		name     string
		args     string
		code     int
		expected string
	}{
		{
			name:     "Testing delete with correct request argument",
			args:     "1",
			code:     200,
			expected: `{"message":"Todo successfully deleted"}`,
		},

		{
			name:     "Testing delete with argument that does not exist within slice",
			args:     "10",
			code:     404,
			expected: `{"error":"Invalid todo ID"}`,
		},

		{
			name:     "Testing delete with incorrect request argument",
			args:     "ab",
			code:     400,
			expected: `{"error":"Invalid URL format"}`,
		},
	}

	for _, test := range tests {
		request := httptest.NewRequest(http.MethodDelete, "/todo/"+test.args, nil)
		response := httptest.NewRecorder()

		DeleteTodo(response, request)

		responseBody := response.Body.String()
		statusCode := response.Result().StatusCode

		if responseBody != test.expected {
			t.Errorf("got %v, want %v", response.Body.String(), test.expected)
		}

		if statusCode != test.code {
			t.Errorf("got %v, want %v", response.Result().StatusCode, test.code)
		}
	}
}

func TestUpdateTodo(t *testing.T) {

	test := struct {
		name     string
		args     string
		expected string
		code     int
	}{
		name:     "Testing update todo endpoint",
		args:     "1",
		expected: `{"message":"Todo status updated to completed"}`,
		code:     200,
	}

	request := httptest.NewRequest(http.MethodPatch, "/todo/"+test.args, nil)
	response := httptest.NewRecorder()

	UpdateTodo(response, request)

	responseBody := response.Body.String()
	statusCode := response.Result().StatusCode

	if responseBody != test.expected {
		t.Errorf("got %v, want %v", response.Body.String(), test.expected)
	}

	if statusCode != test.code {
		t.Errorf("got %v, want %v", response.Result().StatusCode, test.code)
	}

}
