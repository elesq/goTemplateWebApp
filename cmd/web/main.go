package main

import (
	"log"
	"net/http"

	"github.com/elesq/gotemplatewebapp/pkg/config"
	"github.com/elesq/gotemplatewebapp/pkg/handlers"
	"github.com/elesq/gotemplatewebapp/pkg/render"
)

const portNumber = ":8080"

// main function sets up the HTTP server and routes
func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("error crearing templates cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", repo.Home)
	// http.HandleFunc("/about", repo.About)

	log.Printf("Starting server on port %s\n", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
	// _ = http.ListenAndServe(portNumber, nil)
}
