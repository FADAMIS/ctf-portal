package entities

type Challenge struct {
	Name        string          `json:"name"`
	Files       []ChallengeFile `json:"files"`
	Flag        string          `json:"flag"`
	Points      int             `json:"points"`
	Description string          `json:"description"`
	CountryCode string          `json:"country"`
}

type ChallengeFile struct {
	FileName string `json:"filename"`
	Base64   string `json:"base64"`
}

type Challenges struct {
	Challenges []Challenge `json:"challenges"`
}

// this will be sent to the client, basically challenge struct stripped by the flag
type ReturnChallenge struct {
	Name        string          `json:"name"`
	Files       []ChallengeFile `json:"files"`
	Points      int             `json:"points"`
	Description string          `json:"description"`
	CountryCode string          `json:"country"`
}

type ReturnChallenges struct {
	Challenges []ReturnChallenge `json:"challenges"`
}
