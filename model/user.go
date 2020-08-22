package model

type User struct {
	ID             int64  `json:"id"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"passwordRepeat"`
	Image          string `json:"image"`
}
