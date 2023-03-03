package entities

type Session struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	ExpiresIn int64  `json:"expiresin"`
}

type Sessions struct {
	Sessions []Session `json:"sessions"`
}
