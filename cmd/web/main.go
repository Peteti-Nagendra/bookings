package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Peteti-Nagendra/bookings/internal/config"
	"github.com/Peteti-Nagendra/bookings/internal/handlers"
	"github.com/Peteti-Nagendra/bookings/internal/helpers"
	"github.com/Peteti-Nagendra/bookings/internal/models"
	"github.com/Peteti-Nagendra/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNum = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main function of application
func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
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

func run() error {
	gob.Register(models.Reservations{})

	//This varible has to be changes in production env
	app.InProduction = false
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
		return err

	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	helpers.NewHelper(&app)

	return nil

}
