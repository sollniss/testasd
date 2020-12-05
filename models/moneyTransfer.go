package model

import (
	"encoding/json"
	"strings"
	"time"
)

// MoneyTransfer is a basic representation of a transfer entry
type MoneyTransfer struct {
	ID       int64
	Date     time.Time
	Amount   int
	TypeID   int
	PersonID int
	Comment  string
}

// MoneyTransferJSON holds user data with a YYYY-MM-DD format
type MoneyTransferJSON struct {
	Date     TransferDate `json:"date"`
	Amount   int          `json:"amount"`
	TypeID   int          `json:"type_id"`
	PersonID int          `json:"person_id"`
	Comment  string       `json:"comment"`
}

// TransferDate is a type that holds dates in a YYYY-MM-DD format
type TransferDate time.Time

// UnmarshalJSON interface implementation
func (j *TransferDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = TransferDate(t)
	return nil
}

// MarshalJSON interface implementation
func (j TransferDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}

// Format override
func (j TransferDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}
