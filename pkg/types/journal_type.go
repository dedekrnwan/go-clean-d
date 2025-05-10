package types

import "encoding/json"

type JournalType string

const (
	JournalTypeCredit JournalType = "credit"
	JournalTypeDebit  JournalType = "debit"
)

func NewJournalType(in string) (out JournalType, err error) {
	out = JournalType(in)
	return
}

func (dep JournalType) String() string {
	return string(dep)
}

func (dep JournalType) MarshallJSON() ([]byte, error) {
	return json.Marshal(dep)
}
