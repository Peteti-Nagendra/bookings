package forms

type error map[string][]string

// Add adds an error message for the given field.
func (e error) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns the first error message for the given field. If there are no
func (e error) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
