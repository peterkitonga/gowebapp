package form

import (
	"net/url"
	"time"

	"github.com/cohesivestack/valgo"
)

// Forms creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Valid returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Required checks for required fields
func (f *Form) Required(fields map[string]string) {
	for name, title := range fields {
		value := f.Get(name)
		validate := valgo.Is(
			valgo.String(value, name, title).Not().Blank("The '{{title}}' field cannot be blank."),
		)

		handleValidationError(f, validate)
	}
}

// IsAfterDate checks if value is after the expected date
func (f *Form) IsAfterDate(name, title, date string) {
	layout := "2006-01-02"
	value, _ := time.Parse(layout, f.Get(name))
	parsedDate, _ := time.Parse(layout, date)
	validate := valgo.Is(
		valgo.Time(value, name, title).After(parsedDate, "The '{{title}}' cannot be before '{{value}}'."),
	)

	handleValidationError(f, validate)
}

// handleValidationError handles validation errors
func handleValidationError(f *Form, validate *valgo.Validation) {
	if errInfo := validate.ToValgoError(); errInfo != nil {
		for field, valueErr := range errInfo.Errors() {
			f.Errors.Add(field, valueErr.Messages()[0])
		}
	}
}
