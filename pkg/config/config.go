package config

import (
	"html/template"
	"log"
)

// Application config - adheres to the global rule of not
// importing anything from the app itself to avoid cyclical
// imports but is imported by appkication packages that need
// to reference the config values it holds
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
