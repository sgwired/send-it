package handlers

import (
	"net/http"

	"trinity.jakks.net/Administrator/send-it/cmd/web/config"
	"trinity.jakks.net/Administrator/send-it/cmd/web/pkg/models"
	"trinity.jakks.net/Administrator/send-it/cmd/web/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo ceates a new respository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repositiory for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}

// About is the about page handler
func (m *Repository )About(w http.ResponseWriter, r *http.Request) {
	// perform some logic w/ data
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData {
		StringMap: stringMap,
	})
}
