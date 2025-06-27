package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shreyas1307/go-server/internal/handlers"
	"github.com/shreyas1307/go-server/internal/store/models"
	"github.com/stretchr/testify/assert"
)

type MockServerStore struct{}

func (m *MockServerStore) AddTodo(ctx context.Context, todo models.Todo) error {
	return nil
}

func (m *MockServerStore) GetTodos(ctx context.Context) (models.TodoList, error) {
	return models.TodoList{}, nil
}

func TestAddTodoSuccessful(t *testing.T) {
	th := &handlers.TodoHandler{
		Store: &MockServerStore{},
	}
	todo := models.Todo{
		ID:          "1",
		Title:       "Test",
		Description: "This is a test todo",
		Completed:   true,
	}

	todoMarshal, err := json.Marshal(todo)
	assert.NoError(t, err)

	reader := bytes.NewReader(todoMarshal)
	request := httptest.NewRequest(http.MethodPost, "/todo/add", reader)
	recorder := httptest.NewRecorder()

	th.AddTodo(recorder, request)

	res := recorder.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err)
	defer res.Body.Close()

	expextedResponse := map[string]interface{}{
		"successful": true,
		"message":    "new todo added",
	}
	expectedResponseMarshall, err := json.Marshal(expextedResponse)
	assert.NoError(t, err)

	assert.Equal(t, expectedResponseMarshall, b)
}

func TestAddTodoFailure(t *testing.T) {
	th := &handlers.TodoHandler{
		Store: &MockServerStore{},
	}

	badTodo := map[string]interface{}{
		"yo":   "bad todo lol",
		"alex": "says golang is best",
	}

	badTodoMarshal, err := json.Marshal(badTodo)
	assert.NoError(t, err)

	reader := bytes.NewReader(badTodoMarshal)
	request := httptest.NewRequest(http.MethodPost, "/todo/add", reader)
	recorder := httptest.NewRecorder()

	th.AddTodo(recorder, request)

	res := recorder.Result()
	// below test fails... may have to edit handler code to
	// check if I'm getting the corect request body
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}
