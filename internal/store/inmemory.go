package store

import (
	"context"

	"github.com/shreyas1307/go-server/internal/store/models"
)

type InMemoryStore struct {
	todos []models.Todo
}

// NewInMemoryStore initialises a new in-memory store.
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		todos: []models.Todo{},
	}
}

// AddTodo adds a new todo item to the in-memory store.
func (s *InMemoryStore) AddTodo(ctx context.Context, todos models.Todo) error {
	s.todos = append(s.todos, todos)
	return nil
}

// GetTodos retrieves all todo items from the in-memory store.
// If there are no todos, it returns an empty slice.
func (s *InMemoryStore) GetTodos(ctx context.Context) ([]models.Todo, error) {
	if len(s.todos) == 0 {
		return []models.Todo{}, nil
	}
	return s.todos, nil
}
