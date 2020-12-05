package model

// Person is a basic representation of a person
type Person struct {
	ID           int
	Name         string
	Abbreviation string
}

// PersonDict is a wrapper for easy access by ID and looping though the Persons by insertion order.
type PersonDict struct {
	persons map[int]*Person
	Ordered []int
}

// NewPersonDict creates a new PersonDict.
func NewPersonDict() *PersonDict {
	var pd PersonDict
	pd.persons = make(map[int]*Person)
	pd.Ordered = make([]int, 0)
	return &pd
}

// Add inserts the element into the map and appends it to the Ordered slice.
func (pd *PersonDict) Add(p *Person) {
	pd.persons[p.ID] = p
	pd.Ordered = append(pd.Ordered, p.ID)
}

// Get returns a person with the corresponding ID.
func (pd *PersonDict) Get(id int) *Person {
	return pd.persons[id]
}

// Contains if the underlaying map contains a Person with the given ID.
func (pd *PersonDict) Contains(id int) bool {
	_, ok := pd.persons[id]
	return ok
}
