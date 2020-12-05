package model

// TransferType is a basic representation of a money transfer type.
type TransferType struct {
	ID   int
	Name string
}

// TransferTypeDict is a wrapper for easy access by ID and looping though the TransferTypes by insertion order.
type TransferTypeDict struct {
	transferTypes map[int]*TransferType
	Ordered       []int
}

// NewTransferTypeDict creates a new TransferTypeDict.
func NewTransferTypeDict() *TransferTypeDict {
	var tt TransferTypeDict
	tt.transferTypes = make(map[int]*TransferType)
	tt.Ordered = make([]int, 0)
	return &tt
}

// Add inserts the element into the map and appends it to the Ordered slice.
func (ttd *TransferTypeDict) Add(tt *TransferType) {
	ttd.transferTypes[tt.ID] = tt
	ttd.Ordered = append(ttd.Ordered, tt.ID)
}

// Get returns a transfer type with the corresponding ID.
func (ttd *TransferTypeDict) Get(id int) *TransferType {
	return ttd.transferTypes[id]
}

// Contains if the underlaying map contains a TransferType with the given ID.
func (ttd *TransferTypeDict) Contains(id int) bool {
	_, ok := ttd.transferTypes[id]
	return ok
}
