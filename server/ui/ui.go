package ui

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var templates *template.Template

func ParseTemplates() {
	var allFiles []string
	files, err := ioutil.ReadDir("frontend/html")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".html") {
			allFiles = append(allFiles, "frontend/html/"+filename)
		}
	}
	templates, err = template.ParseFiles(allFiles...)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t := templates.Lookup("main.html")

	renderTemplate(t, w, nil)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	t := templates.Lookup("users.html")

	renderTemplate(t, w, nil)
}

func renderTemplate(t *template.Template, w http.ResponseWriter, data interface{}) {
	if t != nil {
		e := t.Execute(w, data)
		if e != nil {
			log.Printf("error rendering template: %s with error : %s", t.Name(), e.Error())
		}
	}
}
