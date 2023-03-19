package entities

type Announcement struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
}

type Announcements struct {
	Announcements []Announcement `json:"announcements"`
}
