package model

import "time"

// Income is a basic representation of an income entry
type Income struct {
	ID       int
	Date     time.Time
	Amount   int
	TypeID   int
	PersonID int
	Comment  string
}

// NewIncome creates a new income entry
func NewIncome(id int, date time.Time, amount int, incomeTypeID int, personID int, comment string) *Income {
	return &Income{
		ID:       id,
		Date:     date,
		Amount:   amount,
		TypeID:   incomeTypeID,
		PersonID: personID,
		Comment:  comment,
	}
}

// Update updates all fields an income entry besides the ID
func (r *Income) Update(date time.Time, amount int, incomeTypeID int, personID int, comment string) {
	r.Date = date
	r.Amount = amount
	r.TypeID = incomeTypeID
	r.PersonID = personID
	r.Comment = comment
}

// GetID returns the database ID for an income entry
func (r *Income) GetID() int {
	return r.ID
}

// GetDate returns the date of an income entry
func (r *Income) GetDate() time.Time {
	return r.Date
}

// GetAmount returns the amount of an income entry
func (r *Income) GetAmount() int {
	return r.Amount
}

// GetTypeID returns the type id for an income entry
func (r *Income) GetTypeID() int {
	return r.TypeID
}

// GetPersonID returns the person id for an income entry
func (r *Income) GetPersonID() int {
	return r.PersonID
}

// GetComment returns the comment for an income entry
func (r *Income) GetComment() string {
	return r.Comment
}
