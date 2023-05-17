package entities

type TeamPoints struct {
	Team        string `json:"team"`
	PointAmount int    `json:"points"`
	Solved      string `json:"solved"`
}

type AllPoints struct {
	Points []TeamPoints `json:"allpoints"`
}
