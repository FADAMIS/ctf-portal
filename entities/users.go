package entities

type User struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Session  Session `json:"session"`
}

type Users struct {
	Users []User `json:"users"`
}
