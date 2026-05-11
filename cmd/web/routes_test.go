package main

import (
	"fmt"
	"testing"

	"github.com/Peteti-Nagendra/bookings/internal/config"
	"github.com/go-chi/chi/v5"
)

func TestRuotes(t *testing.T) {
	var app config.AppConfig
	mux := Routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:

	default:
		t.Errorf(fmt.Sprintf("type is not *chi.mux, but type %v", v))
	}
}
