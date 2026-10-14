package handlers

import (
	"log"
	"net/http"

	"github.com/peterkitonga/gowebapp/internal/config"
	"github.com/peterkitonga/gowebapp/internal/models"
	"github.com/peterkitonga/gowebapp/internal/render"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo the repository used by the handlers
var Repo *Repository

// NewRepo create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["name"] = "John Doe"

	err := render.RenderPage(w, "home.page.jet", &models.TemplateData{
		Data: data,
	})
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!"

	err := render.RenderPage(w, "about.page.jet", &models.TemplateData{
		StringMap: stringMap,
	})
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

// Contact is the contact page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	err := render.RenderPage(w, "contact.page.jet", &models.TemplateData{})
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

// Reservation is the reservation page handler
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	err := render.RenderPage(w, "reservation.page.jet", &models.TemplateData{})
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

// Quarters is the general's quarters page handler
func (m *Repository) Quarters(w http.ResponseWriter, r *http.Request) {
	err := render.RenderPage(w, "room-generals-quarters.page.jet", &models.TemplateData{})
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

// Suite is the major's suite page handler
func (m *Repository) Suite(w http.ResponseWriter, r *http.Request) {
	err := render.RenderPage(w, "room-majors-suite.page.jet", &models.TemplateData{})
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
