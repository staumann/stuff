package ui

import (
	"fmt"
	"github.com/staumann/caluclation/database"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	billRepository database.BillRepository
	userRepository database.UserRepository
	templates      *template.Template
)

func Prepare(b database.BillRepository, u database.UserRepository) {
	billRepository = b
	userRepository = u
}

func ParseTemplates(path string) {
	var allFiles []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".html") {
			log.Printf("Parsing Template: %s", filename)
			allFiles = append(allFiles, fmt.Sprintf("%s/%s", path, filename))
		}
	}
	templates, err = template.ParseFiles(allFiles...)
	if err != nil {
		log.Printf("error parsing templates: %v", err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t := templates.Lookup("main.html")

	renderTemplate(t, w, nil)
}

func BillHandler(w http.ResponseWriter, r *http.Request) {
	t := templates.Lookup("bills.html")
	if t == nil {
		log.Print("FUCK")
	}
	renderTemplate(t, w, make(map[string]interface{}))
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(http.StatusInternalServerError)
	t := templates.Lookup("error.html")
	if t != nil {
		e := t.Execute(w, map[string]string{"error": err.Error()})
		if e != nil {
			log.Printf("error serving error page: %v", e)
		}
	}
}

func renderTemplate(t *template.Template, w http.ResponseWriter, data interface{}) {
	if t != nil {
		e := t.Execute(w, data)
		if e != nil {
			log.Printf("error rendering template: %s with error : %s", t.Name(), e.Error())
		}
	}
}
