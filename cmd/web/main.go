package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/peterkitonga/gowebapp/internal/config"
	"github.com/peterkitonga/gowebapp/internal/handler"
	"github.com/peterkitonga/gowebapp/internal/render"
)

const portNumber = ":8180"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port%s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	err := srv.ListenAndServe()
	log.Fatal(err)
}
