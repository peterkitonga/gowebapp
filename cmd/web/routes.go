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

	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)
	mux.Get("/contact", handler.Repo.Contact)

	mux.Get("/make-reservation", handler.Repo.Reservation)

	mux.Get("/rooms/generals-quarters", handler.Repo.Quarters)
	mux.Get("/rooms/majors-suite", handler.Repo.Suite)

	mux.Get("/search-availability", handler.Repo.SearchAvailability)
	mux.Post("/search-availability", handler.Repo.PostSearchAvailability)

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
