package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/peterkitonga/gowebapp/internal/config"
	"github.com/peterkitonga/gowebapp/internal/handlers"
	"github.com/peterkitonga/gowebapp/internal/render"
)

const portNumber = ":8180"

func main() {
	var app config.AppConfig

	app.InProduction = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port%s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	err := srv.ListenAndServe()
	log.Fatal(err)
}
