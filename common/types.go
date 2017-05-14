package common

type LogEntry struct {
	Body        string `json:"body"`
	DateCreated string `json:"dateCreated"`
}

type Ticket struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type Project struct {
	Name    string   `json:"name"`
	Tickets []Ticket `json:"tickets"`
}
