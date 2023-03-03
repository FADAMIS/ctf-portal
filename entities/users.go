package entities

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Users struct {
	Users []User `json:"users"`
}
