package types

import "encoding/json"

type Currency string

const (
	CurrencyIDR = "IDR"
)

func (dep Currency) String() string {
	return string(dep)
}

func (dep Currency) MarshallJSON() ([]byte, error) {
	return json.Marshal(dep)
}
