package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/peterkitonga/gowebapp/internal/config"
	"github.com/peterkitonga/gowebapp/internal/model"
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

	err := render.RenderPage(w, r, "home.page.jet", &models.TemplateData{
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

	err := render.RenderPage(w, r, "about.page.jet", &models.TemplateData{
		StringMap: stringMap,
	})
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

// Contact is the contact page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	err := render.RenderPage(w, r, "contact.page.jet", &models.TemplateData{})
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

// Reservation is the reservation page handler
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	err := render.RenderPage(w, r, "reservation.page.jet", &models.TemplateData{})
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

// Quarters is the general's quarters page handler
func (m *Repository) Quarters(w http.ResponseWriter, r *http.Request) {
	err := render.RenderPage(w, r, "room-generals-quarters.page.jet", &models.TemplateData{})
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

// Suite is the major's suite page handler
func (m *Repository) Suite(w http.ResponseWriter, r *http.Request) {
	err := render.RenderPage(w, r, "room-majors-suite.page.jet", &models.TemplateData{})
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

// SearchAvailability is the search availability page handler
func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	err := render.RenderPage(w, r, "search-availability.page.jet", &models.TemplateData{})
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

// PostSearchAvailability handles request for search availability
func (m *Repository) PostSearchAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	m.App.Session.Put(r.Context(), "flash", fmt.Sprintf("Rooms available from '%s' to '%s'!", start, end))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
