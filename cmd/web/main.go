package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"

	"github.com/rtmelsov/mansProducts/pkg/config"
	"github.com/rtmelsov/mansProducts/pkg/handlers"
	"github.com/rtmelsov/mansProducts/pkg/render"
)

// port number to work on develop
const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application
func main() {
	// change this to true when in production
	app.InProduction = false

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	app.TemplateCache = tc
	app.UseCache = true

	render.NewTemplates(&app)

	repo := handlers.GetNewRepository(&app)
	handlers.NewHandler(repo)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprint("Starting the application on port", portNumber))

	// _ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
