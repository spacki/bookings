package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/spacki/bookings/pkg/config"
	"github.com/spacki/bookings/pkg/handlers"
	"github.com/spacki/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8088"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	app.UseCache = false
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction


	app.Session = session

	app.UseCache = true
	var err error
	app.TemplateCache, err = render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache, %s", err)
	}

	repo := handlers.NewRepo(&app)
	handlers.Newhandlers(repo)

	render.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)
	log.Printf( "starting application on port %s \n", portNumber)
	//http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}
