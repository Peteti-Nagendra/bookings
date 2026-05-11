package models

import "github.com/Peteti-Nagendra/bookings/internal/forms"

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	floatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
