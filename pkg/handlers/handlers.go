package handlers

import (
	"net/http"
	"github.com/DarioAcevedo/go-course/pkg/render"
)


func Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about")
}

