package handlers

import (
	"net/http"

	"github.com/peterkitonga/gowebapp/pkg/config"
	"github.com/peterkitonga/gowebapp/pkg/models"
	"github.com/peterkitonga/gowebapp/pkg/render"
)

// Repository: is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo: the repository used by the handlers
var Repo *Repository

// NewRepo: create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers: sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home: is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About: is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!"

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
