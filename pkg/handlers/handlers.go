package handlers

import (
	"github.com/spacki/bookings/pkg/config"
	"github.com/spacki/bookings/pkg/models"
	"github.com/spacki/bookings/pkg/render"
	"log"
	"net/http"
	"strings"
)


var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func Newhandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home handler called ")
	log.Printf("using cache %t", m.App.UseCache)
	stringMap := make(map[string]string)
	stringMap["test"] = "Eisern Union!"
	remoteSocket := strings.Split(r.RemoteAddr, ":")

	log.Printf("request fired from IP-address %s and port %s", remoteSocket[0], remoteSocket[1])
	remoteIP := remoteSocket[0]
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	myTest := m.App.Session.GetString(r.Context(), "remote_ip" )
	log.Printf("===> quick test %s", myTest)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{StringMap: stringMap,})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	log.Println("about handler called")
	log.Printf("using cache %t", m.App.UseCache)

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	log.Printf("session remote ip address %s", remoteIP)


	// some business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Eisern Union!"

	stringMap["remote_ip"] = remoteIP

	val, ok := stringMap["test"]

	if !ok {
		log.Println("stringMap has no value for key test")
	} else {
		log.Printf("stringmaps value for key test %s \n",val)
	}



	// send it to the page

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap,})
}
