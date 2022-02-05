package main

import (
	"net/http"
	"github.com/DarioAcevedo/go-course/pkg/handlers"
)

const portNumer = ":8080"


func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about",handlers.About)
	_ = http.ListenAndServe(portNumer, nil)
}

