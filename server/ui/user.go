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
		LastName:       r.Form.Get("lastName"),
		FirstName:      r.Form.Get("firstName"),
		Image:          r.Form.Get("image"),
		Password:       r.Form.Get("password"),
		PasswordRepeat: r.Form.Get("passwordRepeat"),
	}

	errs := checkNewUser(user)
	if len(errs) > 0 {
		t := templates.Lookup("users_new.html")
		w.WriteHeader(http.StatusOK)
		_ = t.Execute(w, map[string]interface{}{
			"FirstName":      user.FirstName,
			"LastName":       user.LastName,
			"Image":          user.Image,
			"Password":       user.Password,
			"PasswordRepeat": user.PasswordRepeat,
			"Errors":         errs,
		})

		return
	}

	if e := userRepository.SaveUser(user); e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error storing user: %v", e)
	} else {
		http.Redirect(w, r, "/users", http.StatusMovedPermanently)
	}
}

func checkNewUser(user *model.User) []string {
	errors := make([]string, 0)

	if user.FirstName == "" {
		errors = append(errors, "Firstname was missing")
	}

	if user.LastName == "" {
		errors = append(errors, "Lastname was missing")
	}
	if user.Password == "" {
		errors = append(errors, "Password is missing")
	}
	if user.PasswordRepeat != user.Password {
		errors = append(errors, "Passowords were not identical")
	}
	return errors
}
