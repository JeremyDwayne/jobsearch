package main

import (
	view "jobsearch/app/views"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("./public/assets"))
	r.Handle("/public/assets/*", http.StripPrefix("/public/assets/", fs))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		view.Index().Render(r.Context(), w)
	})
	http.ListenAndServe(":3000", r)
}
