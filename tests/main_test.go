package main_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, "Hello world", string(body))
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

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, "200 - OK", string(body))
}
