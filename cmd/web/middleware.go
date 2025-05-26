package main

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

// WriteToConsole is an example middleware function which serves
// no useful purpose other than to demonstrate the construction
// of a hand-cooked middleware function
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

// NoSurf is a middleware function that adds CSRF protection
// to the application. It uses the nosurf package to create
// a CSRF token and set it in a cookie.
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.IsProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad is a middleware function that loads the session
// and saves it after the request is processed. It uses the
// session manager from the app config to manage sessions.
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
