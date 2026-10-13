package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/peterkitonga/gowebapp/internal/handlers"
)

func routes() http.Handler {
	fileServer := http.FileServer(http.Dir("./web/static"))
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
