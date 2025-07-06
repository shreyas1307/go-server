package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/shreyas1307/go-server/internal/store"
	"github.com/shreyas1307/go-server/internal/store/models"
)

// TodoHandler handles requests related to the todo resource.
type TodoHandler struct {
	Store store.Store
}

// todoResponse is the structure of the response for todo-related requests.
type todoResponse struct {
	Successful bool          `json:"successful"`
	Message    string        `json:"message"`
	Data       []models.Todo `json:"data,omitempty"`
	Error      error         `json:"error,omitempty"`
}

// Router sets up the routes for the TodoHandler.
func (th *TodoHandler) Router() chi.Router {
	router := chi.NewRouter()
	router.Get("/", th.GetTodos)
	router.Post("/add", th.AddTodo)
	return router
}

// NewTodoHandler creates a new TodoHandler with the provided store.
//
// It returns an error if the store is nil.
func NewTodoHandler(store store.Store) (*TodoHandler, error) {
	if store == nil {
		return nil, errors.New("store cannot be nil")
	}
	return &TodoHandler{
		Store: store,
	}, nil
}

// GetTodos handles the request to fetch all todo items.
func (th *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	todoItems, err := th.Store.GetTodos(ctx)
	println(todoItems) // Debugging, I'm unable to print the todo items for the response
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		res, _ := json.Marshal(&todoResponse{
			Successful: false,
			Message:    "Error fetching todos",
			Error:      err,
		})
		w.Write(res)
		return
	}
	w.WriteHeader(http.StatusOK)
	res, err := json.Marshal(&todoResponse{
		Successful: true,
		Message:    "Todos fetched successfully",
		Data:       todoItems,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// AddTodo handles the request to add a new todo item.
func (th *TodoHandler) AddTodo(w http.ResponseWriter, r *http.Request) {

	var todo models.Todo
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(&todoResponse{
			Successful: false,
			Message:    "Invalid request body",
			Error:      err,
		})
		w.Write(res)
		return
	}

	res, err := json.Marshal(&todoResponse{
		Successful: true,
		Message:    "new todo added",
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
