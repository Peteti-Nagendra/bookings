package main

import (
	"net/http"

	"github.com/Peteti-Nagendra/bookings/pkg/config"
	"github.com/Peteti-Nagendra/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func Routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewMux()
	// mux.Use(middleware.Recoverer)
	// mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
