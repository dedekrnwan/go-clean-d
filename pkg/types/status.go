package types

import "encoding/json"

type StatusTransaction string

const (
	StatusInitiated StatusTransaction = "initiated"
	StatusPending   StatusTransaction = "pending"
	StatusSuccess   StatusTransaction = "success"
	StatusFailed    StatusTransaction = "failed"
	StatusExpired   StatusTransaction = "expired"
	StatusRefunded  StatusTransaction = "refunded"
	StatusCancelled StatusTransaction = "cancelled"
)

func (dep StatusTransaction) String() string {
	return string(dep)
}

func (dep StatusTransaction) MarshallJSON() ([]byte, error) {
	return json.Marshal(dep)
}

type StatusAccount string

const (
	StatusActive   StatusAccount = "active"
	StatusFreeze   StatusAccount = "freeze"
	StatusClosed   StatusAccount = "closed"
	StatusDormant  StatusAccount = "dormant"
	StatusInactive StatusAccount = "inactive"
)

func (dep StatusAccount) String() string {
	return string(dep)
}

func (dep StatusAccount) MarshallJSON() ([]byte, error) {
	return json.Marshal(dep)
}
