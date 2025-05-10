package types

import "encoding/json"

type BalanceType string

const (
	BalanceTypeLedger BalanceType = "ledger"
	BalanceTypeFloat  BalanceType = "float"
	BalanceTypeHold   BalanceType = "hold"
)

func (dep BalanceType) String() string {
	return string(dep)
}

func (dep BalanceType) MarshallJSON() ([]byte, error) {
	return json.Marshal(dep)
}
