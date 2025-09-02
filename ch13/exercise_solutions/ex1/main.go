package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC3339)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(t))
}

func router() chi.Router {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		getRoot(w, r)
	})

	return r
}

func main() {
	r := router()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
