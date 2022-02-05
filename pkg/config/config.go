package config

import "html/template"


// HOlds the application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache bool
}