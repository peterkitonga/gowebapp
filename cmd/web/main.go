package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/peterkitonga/gowebapp/pkg/config"
	"github.com/peterkitonga/gowebapp/pkg/handlers"
	"github.com/peterkitonga/gowebapp/pkg/render"
)

const portNumber = ":8180"

func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port%s\n", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}
