package main_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/hello-world")
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

	body, err := io.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, "Hello world", string(body))
}

func TestTodoHandler(t *testing.T) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/todo", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		w.Write([]byte("Received todo: " + string(body)))
	})

	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Post(ts.URL+"/todo", "application/json", io.NopCloser(strings.NewReader(`{"title":"Test Todo"}`)))
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

	body, err := io.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, "Received todo: {\"title\":\"Test Todo\"}", string(body))
}

func TestHealthz(t *testing.T) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("200 - OK"))
	})

	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/healthz")
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

	body, err := io.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, "200 - OK", string(body))
}
