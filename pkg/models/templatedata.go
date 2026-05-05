package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	floatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	error     string
}
