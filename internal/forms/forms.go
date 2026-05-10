package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Forms creates a new form
type Forms struct {
	url.Values
	Errors error
}

// New creates a new form
func New(data url.Values) *Forms {
	return &Forms{
		data,
		error(map[string][]string{}),
	}
}

// Has checks if the given field is present in the request
func (f *Forms) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}

// Valid checks if the form is valid
func (f *Forms) Valid() bool {
	return len(f.Errors) == 0
}

// Required checks if the given field is present in the request
func (f *Forms) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {

			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Minlength checks if the given field is present in the request
func (f *Forms) MinLen(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

func (f *Forms) EmailValidator(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
