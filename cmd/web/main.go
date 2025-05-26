package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/elesq/gotemplatewebapp/pkg/config"
	"github.com/elesq/gotemplatewebapp/pkg/handlers"
	"github.com/elesq/gotemplatewebapp/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main function sets up the HTTP server and routes
func main() {

	app.IsProduction = false // set to true in production

	// define the session and add it to the app config
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.IsProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("error crearing templates cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	log.Printf("Starting server on port %s\n", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
