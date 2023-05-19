package entities

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ID int `json:"id"`
}

type Users struct {
	Users []User `json:"users"`
}
