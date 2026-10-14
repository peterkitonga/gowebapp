package render

import (
	"log"
	"net/http"
	"reflect"

	"github.com/CloudyKit/jet/v6"
	"github.com/peterkitonga/gowebapp/internal/config"
	"github.com/peterkitonga/gowebapp/internal/model"
)

var app *config.AppConfig

// NewTemplates sets the config for the template function
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData set the default data passed to templates
func AddDefaultData(templateData *models.TemplateData) *models.TemplateData {
	return templateData
}

// RenderPage renders templates using jet templates
func RenderPage(w http.ResponseWriter, tmpl string, td *models.TemplateData) error {
	var templateSet = jet.NewSet(
		jet.NewOSFileSystemLoader("./web/templates"),
		jet.DevelopmentMode(app.InProduction),
	)
	td = AddDefaultData(td)
	data := DataToVarMap(td)

	template, err := templateSet.GetTemplate(tmpl)
	if err != nil {
		log.Fatalf("Error loading template: %v", err)

		return err
	}

	err = template.Execute(w, data, nil)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)

		return err
	}

	return nil
}

func DataToVarMap(td *models.TemplateData) jet.VarMap {
	data := make(jet.VarMap)
	v := reflect.ValueOf(&td)
	// Resolves to the underlying struct
	v = reflect.Indirect(v)

	// Handle pointers gracefully
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Only process structs
	if v.Kind() != reflect.Kind(reflect.Struct) {
		return data
	}

	// Get field names and metadata
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)

		// Skip unexported fields
		if field.PkgPath != "" {
			continue
		}

		// Use the field name as the variable key
		data.Set(field.Name, v.Field(i).Interface())
	}

	return data
}
