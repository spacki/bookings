package render

import (
	"bytes"
	"fmt"
	"github.com/spacki/bookings/pkg/config"
	"github.com/spacki/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions template.FuncMap

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}


// adding something we need for every page
func addDefaultData(td *models.TemplateData, title string) *models.TemplateData {
	td.StringMap["title"] = title
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	log.Printf("render template %s \n", tmpl)
	log.Printf("using cache %t", app.UseCache)
	// getting the templates from the config
	var templateCache map[string]*template.Template
	// in production use cache, for development build the templates new for each call
	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("something went wrong")
	}


	td = addDefaultData(td, tmpl)

	buf := new(bytes.Buffer)
	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to the browser")
	}
}

// create a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
     log.Println("create a template cache only called once during application startup")
	// create a map the keys are the name of the original not parsed temples
	// in our case there are two: about.page.tmpl and home.page.tmpl
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	log.Printf("we have %d templates to process \n", len(pages))
	// we are processing two pages
	for _, page := range pages {
		name := filepath.Base(page)
		log.Printf("process template %s \n", name)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// now looking for layout files
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		log.Printf("we found %d layout files \n", len(matches))
		if len(matches) > 0 {
			log.Println("combining layouts with templates")
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
