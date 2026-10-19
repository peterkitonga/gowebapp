package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/peterkitonga/gowebapp/internal/config"
	"github.com/peterkitonga/gowebapp/internal/form"
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
	err := render.RenderPage(w, r, "home.page.jet", &model.TemplateData{})
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!"

	err := render.RenderPage(w, r, "about.page.jet", &model.TemplateData{
		StringMap: stringMap,
	})
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

// Contact is the contact page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	err := render.RenderPage(w, r, "contact.page.jet", &model.TemplateData{})
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

// Reservation is the reservation page handler
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyForm model.ReservationForm

	data := make(map[string]interface{})
	data["reservation"] = emptyForm

	err := render.RenderPage(w, r, "reservation.page.jet", &model.TemplateData{
		Form: form.New(nil),
		Data: data,
	})
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

// PostReservation handles the request for reservation
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	fieldLabels := make(map[string]string)
	fieldLabels["first_name"] = "First Name"
	fieldLabels["last_name"] = "Last Name"
	fieldLabels["email"] = "Email"
	fieldLabels["phone"] = "Phone"

	reservation := model.ReservationForm{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := form.New(r.PostForm)
	form.Required(fieldLabels)
	form.HasMinLength("first_name", "First Name", 3)
	form.HasMinLength("last_name", "Last Name", 3)
	form.IsEmail("email", "Email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderPage(w, r, "reservation.page.jet", &model.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

// ReservationSummary is the reservation summary page handler
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(model.ReservationForm)
	if !ok {
		log.Printf("Cannot get item from session")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation

	err := render.RenderPage(w, r, "reservation-summary.page.jet", &model.TemplateData{
		Data: data,
	})
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

// Quarters is the general's quarters page handler
func (m *Repository) Quarters(w http.ResponseWriter, r *http.Request) {
	err := render.RenderPage(w, r, "room-generals-quarters.page.jet", &model.TemplateData{})
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

// Suite is the major's suite page handler
func (m *Repository) Suite(w http.ResponseWriter, r *http.Request) {
	err := render.RenderPage(w, r, "room-majors-suite.page.jet", &model.TemplateData{})
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

// SearchAvailability is the search availability page handler
func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	var emptyForm model.AvailabilityForm

	data := make(map[string]interface{})
	data["availability"] = emptyForm

	err := render.RenderPage(w, r, "search-availability.page.jet", &model.TemplateData{
		Form: form.New(nil),
		Data: data,
	})
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

// PostSearchAvailability handles request for search availability
func (m *Repository) PostSearchAvailability(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	fieldLabels := make(map[string]string)
	fieldLabels["start"] = "Starting Date"
	fieldLabels["end"] = "Ending Date"

	availability := model.AvailabilityForm{
		StartDate: r.Form.Get("start"),
		EndDate:   r.Form.Get("end"),
	}

	form := form.New(r.PostForm)
	form.Required(fieldLabels)
	form.IsAfterDate("end", "Ending Date", availability.StartDate)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["availability"] = availability

		render.RenderPage(w, r, "search-availability.page.jet", &model.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "flash", fmt.Sprintf("Rooms available from '%s' to '%s'!", availability.StartDate, availability.EndDate))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
