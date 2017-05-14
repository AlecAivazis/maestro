package common

import (
	"encoding/base64"
	"fmt"
	"strings"
)

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

type GlobalID struct {
	Type string
	ID   string
}

// GlobalDelim is the string that joins the ID and Type of a GlobalID.
var GlobalDelim = "::"

// MarshalGlobalID returns a globally unique ID for every entity in maestro.
func MarshalGlobalID(ID *GlobalID) string {
	// the string representing the ID
	str := []byte(strings.Join([]string{ID.Type, ID.ID}, GlobalDelim))

	// return the base64 encoded version of the string
	return base64.StdEncoding.EncodeToString(str)
}

func UnmarshalGlobalID(ID string) (*GlobalID, error) {
	// decode the string
	str, err := base64.StdEncoding.DecodeString(ID)
	if err != nil {
		return nil, err
	}
	// treat the
	idInfo := strings.Split(string(str), GlobalDelim)

	// if there are more than two entries
	if len(idInfo) != 2 {
		return nil, fmt.Errorf("ID has too many instances of '%v'", GlobalDelim)
	}

	// return a global id with the
	return &GlobalID{
		Type: idInfo[0],
		ID:   idInfo[1],
	}, nil
}
