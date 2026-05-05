package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf is a middleware function that adds CSRF protection to the application.
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad is a middleware function that loads the session from the request and saves it back to the response.
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}