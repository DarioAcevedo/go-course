package render

import (
	"fmt"
	"html/template"
	"net/http"
)



func RenderTemplate(w http.ResponseWriter, tmp string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmp + ".gohtml")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}