package main

import (
	"github.com/justinas/nosurf"
	"log"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("hit a page")
		next.ServeHTTP(w,r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	log.Println("csrf nosurf middleware called")
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
    return csrfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	log.Printf("++sessionload called++")
	return session.LoadAndSave(next)
}