package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/peterkitonga/gowebapp/pkg/config"
	"github.com/peterkitonga/gowebapp/pkg/models"
)

var app *config.AppConfig

// NewTemplates: sets the config for the template function
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData: set the default data passed to templates
func AddDefaultData(templateData *models.TemplateData) *models.TemplateData {
	return templateData
}

// CreateTemplateCache: creates a cache for the templates
func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := make(map[string]*template.Template)

	// get all of the files named *.page.html from ./web/templates
	pages, err := filepath.Glob("./web/templates/*.page.html")
	if err != nil {
		return templateCache, err
	}

	// range through all files ending with *.page.html
	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		matches, err := filepath.Glob("./web/templates/*.layout.html")
		if err != nil {
			return templateCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./web/templates/*.layout.html")
			if err != nil {
				return templateCache, err
			}
		}

		templateCache[name] = templateSet
	}

	return templateCache, nil
}

// RenderTemplate: renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// create a template cache
	var templateCache map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		templateCache = app.TemplateCache
	} else {
		// otherwise create a new template cache
		templateCache, _ = CreateTemplateCache()
	}

	// get requested template from cache
	template, hasTemplate := templateCache[tmpl]
	if !hasTemplate {
		log.Fatal("Could not get template from template cache")
	}

	// // create a buffer for fine tuned error handling
	buff := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := template.Execute(buff, td)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buff.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}
}
