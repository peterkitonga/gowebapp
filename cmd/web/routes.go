package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/peterkitonga/gowebapp/internal/handler"
)

func routes() http.Handler {
	fileServer := http.FileServer(http.Dir("./web/static"))
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/make-reservation", handlers.Repo.Reservation)

	mux.Get("/rooms/generals-quarters", handlers.Repo.Quarters)
	mux.Get("/rooms/majors-suite", handlers.Repo.Suite)

	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	mux.Post("/search-availability", handlers.Repo.PostSearchAvailability)

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
