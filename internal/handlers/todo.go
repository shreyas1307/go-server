package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/shreyas1307/go-server/internal/store"
	"github.com/shreyas1307/go-server/internal/store/models"
)

type TodoHandler struct {
	Store store.ServerStore
}

func (th *TodoHandler) Router() chi.Router {
	router := chi.NewRouter()
	router.Get("/", th.GetTodos)
	router.Post("/add", th.AddTodo)
	return router
}

func (th *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_, err := th.Store.GetTodos(ctx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Error fetching todos")
		return
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Get Todos \n")
}

func (th *TodoHandler) AddTodo(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Error reading request body \n")
		return
	}
	var todo models.Todo
	if err := json.Unmarshal(b, &todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Error unmarshalling todo")
		return
	}

	res, err := json.Marshal(map[string]interface{}{
		"successful": true,
		"message":    "new todo added",
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
