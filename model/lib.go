package model

type User struct {
	Id        string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Password1 string `json:"password1,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Roles     string `json:"roles,omitempty"`
	Email     string `json:"email,omitempty"`
	Contact   string `json:"contact,omitempty"`
}
