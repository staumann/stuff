package ui

import (
	"github.com/staumann/caluclation/model"
	"log"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	t := templates.Lookup("users.html")
	data := map[string]interface{}{
		"Users": userRepository.GetUsers(),
	}
	renderTemplate(t, w, data)
}

func NewUserHandler(w http.ResponseWriter, r *http.Request) {
	t := templates.Lookup("users_new.html")

	renderTemplate(t, w, nil)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if e := r.ParseForm(); e != nil {
		log.Printf("error parsing form: %v", e)
		handleError(w, r, e)
		return
	}
	user := &model.User{
		LastName:  r.Form.Get("lastName"),
		FirstName: r.Form.Get("firstName"),
		Image:     r.Form.Get("image"),
		Password:  r.Form.Get("password"),
	}
	if e := userRepository.SaveUser(user); e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error storing user: %v", e)
	} else {
		http.Redirect(w, r, "/users", http.StatusMovedPermanently)
	}
}
