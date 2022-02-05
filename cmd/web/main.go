package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/DarioAcevedo/go-course/pkg/config"
	"github.com/DarioAcevedo/go-course/pkg/handlers"
	"github.com/DarioAcevedo/go-course/pkg/render"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	srv := &http.Server{
		Addr: portNumber,
		Handler: router(&app),
	}
	
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
