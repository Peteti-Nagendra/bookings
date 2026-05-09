package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Peteti-Nagendra/bookings/internal/config"
	"github.com/Peteti-Nagendra/bookings/internal/handlers"
	"github.com/Peteti-Nagendra/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNum = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main function of application
func main() {

	//This varible has to be changes in production env
	app.InProduction = true

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)
	render.NewTemplate(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf(fmt.Sprintf("Starting application on port: %s", portNum))

	// _ = http.ListenAndServe(portNum, nil)
	srv := &http.Server{
		Addr:    portNum,
		Handler: Routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
