package entities

type Challenge struct {
	Name   string          `json:"name"`
	Files  []ChallengeFile `json:"files"`
	Flag   string          `json:"flag"`
	Points int             `json:"points"`
}

type ChallengeFile struct {
	FileName string `json:"filename"`
	Base64   string `json:"base64"`
}
