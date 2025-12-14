package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/mux"
)

func main() {

	routerMux := mux.NewRouter()

	sub := routerMux.PathPrefix("/books").Subrouter()

	sub.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		fmt.Fprintf(w, "You requested for the book: %s\n", vars["name"])
	}).Methods("GET").Host("localhost:8080").Schemes("http")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	/* Chi Router */
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{name}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!! It is time to dive into Go"))
	})

	http.ListenAndServe(":8080", r)
}
