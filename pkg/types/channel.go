package types

import "encoding/json"

type ChannelType string

const (
	ChannelTypeBank    ChannelType = "bank"
	ChannelTypeEwallet ChannelType = "ewallet"
)

func (dep ChannelType) String() string {
	return string(dep)
}

func (dep ChannelType) MarshallJSON() ([]byte, error) {
	return json.Marshal(dep)
}
