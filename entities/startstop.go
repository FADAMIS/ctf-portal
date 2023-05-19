package entities

type AutomaticStart struct {
	StartTime int64 `json:"starttime"`
	StopTime  int64 `json:"stoptime"`
	// false if manual start is used
	IsValid bool `json:"isvalid"`
	ID	int	`json:"id"`
}

type ManualStart struct {
	IsStarted bool `json:"isstarted"`

	// false if automatic start is used
	IsValid bool `json:"isvalid"`
	ID	int	`json:"id"`
}

// this is used to write both of entries to the "database"
type StartStopInfo struct {
	Automatic AutomaticStart `json:"automatic"`
	Manual    ManualStart    `json:"manual"`
}
