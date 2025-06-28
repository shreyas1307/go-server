package store

import (
	"context"

	"github.com/shreyas1307/go-server/internal/store/models"
)

type Store interface {
	AddTodo(ctx context.Context, todos models.Todo) error
	GetTodos(ctx context.Context) (models.TodoList, error)
}

type Server struct {
	store Store
}
