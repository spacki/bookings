package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
)

type AppConfig struct {
	UseCache      bool
	InProduction bool
	TemplateCache map[string]*template.Template
	Session *scs.SessionManager
}
