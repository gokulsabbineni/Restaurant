package model

type User struct {
	Id        int    `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Roles     string `json:"roles,omitempty"`
	Email     string `json:"email,omitempty"`
	Contact   string `json:"contact,omitempty"`
}
