package types

import "encoding/json"

type TrxType string

const (
	TrxTypeTopup        TrxType = "topup"
	TrxTypePayment      TrxType = "payment"
	TrxTypePayout       TrxType = "payout"
	TrxTypeTransfer     TrxType = "transfer"
	TrxTypeDisbursement TrxType = "disbursement"
)

func (dep TrxType) String() string {
	return string(dep)
}

func (dep TrxType) MarshallJSON() ([]byte, error) {
	return json.Marshal(dep)
}

type TrxAmountType string

const (
	TrxAmounTypeFee TrxAmountType = "fee"
)

func (dep TrxAmountType) String() string {
	return string(dep)
}

func (dep TrxAmountType) MarshallJSON() ([]byte, error) {
	return json.Marshal(dep)
}

type TrxAmountCategory string

const (
	TrxAmounTypeCategoryInternalFee TrxAmountCategory = "internal_fee"
	TrxAmounTypeCategoryPartnerFee  TrxAmountCategory = "partner_fee"
)

func (dep TrxAmountCategory) String() string {
	return string(dep)
}

func (dep TrxAmountCategory) MarshallJSON() ([]byte, error) {
	return json.Marshal(dep)
}
