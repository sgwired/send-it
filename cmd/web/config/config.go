package config

import (
	"html/template"
	"log"
)

// AppConfg holds the applicatioin config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
