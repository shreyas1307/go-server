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

type MockTodoResponse struct {
	Successful bool          `json:"successful"`
	Message    string        `json:"message"`
	Data       []models.Todo `json:"data,omitempty"`
	Error      error         `json:"error,omitempty"`
}

func (m *MockServerStore) AddTodo(ctx context.Context, todo models.Todo) error {
	return nil
}

func (m *MockServerStore) GetTodos(ctx context.Context) ([]models.Todo, error) {
	return []models.Todo{}, nil
}

func TestNewTodoHandler(t *testing.T) {
	store := &MockServerStore{}
	th, err := handlers.NewTodoHandler(store)
	assert.NoError(t, err)
	assert.NotNil(t, th)
	assert.Equal(t, store, th.Store)
}

func TestNewHandlerNilStore(t *testing.T) {
	th, err := handlers.NewTodoHandler(nil)
	assert.Error(t, err)
	assert.Nil(t, th)
	assert.Equal(t, "store cannot be nil", err.Error())
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

	r := bytes.NewReader(todoMarshal)
	rq := httptest.NewRequest(http.MethodPost, "/todo/add", r)
	rr := httptest.NewRecorder()

	th.AddTodo(rr, rq)

	res := rr.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err)
	defer res.Body.Close()

	response := &MockTodoResponse{
		Successful: true,
		Message:    "new todo added",
	}
	responseMarshal, err := json.Marshal(response)
	assert.NoError(t, err)

	assert.Equal(t, responseMarshal, b)
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

	r := bytes.NewReader(badTodoMarshal)
	rq := httptest.NewRequest(http.MethodPost, "/todo/add", r)
	rr := httptest.NewRecorder()

	th.AddTodo(rr, rq)

	res := rr.Result()
	// below test fails... may have to edit handler code to
	// check if I'm getting the corect request body
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestGetTodosSuccessful(t *testing.T) {
	th := &handlers.TodoHandler{
		Store: &MockServerStore{},
	}

	rq := httptest.NewRequest(http.MethodGet, "/todo/get", nil)
	rr := httptest.NewRecorder()

	th.GetTodos(rr, rq)

	res := rr.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err)
	defer res.Body.Close()

	response := &MockTodoResponse{
		Successful: true,
		Message:    "Todos fetched successfully",
		Data:       []models.Todo{},
	}
	responseMarshal, err := json.Marshal(response)
	assert.NoError(t, err)

	assert.Equal(t, responseMarshal, b)
}
