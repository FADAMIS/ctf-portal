package entities

type TeamPoints struct {
	Team        string `json:"team"`
	PointAmount int    `json:"points"`
	Solved      string `json:"solved"`
	ID	int	`json:"id"`
}

type AllPoints struct {
	Points []TeamPoints `json:"allpoints"`
}
