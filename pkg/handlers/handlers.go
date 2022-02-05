package handlers

import (
	"net/http"

	"github.com/DarioAcevedo/go-course/pkg/config"
	"github.com/DarioAcevedo/go-course/pkg/models"
	"github.com/DarioAcevedo/go-course/pkg/render"
)


type Repository struct {
	App *config.AppConfig
}




// Repository used by the handlers
var Repo *Repository

// Creates a repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

//Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (re *Repository)Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (re *Repository) About(w http.ResponseWriter, r *http.Request) {
	sm := make(map[string]string) 
	sm["damn"] = "I'm data, bitch!"
	td := models.TemplateData{
		StringMap: sm,
	}
	render.RenderTemplate(w, "about.page.tmpl", &td)
}


