package types

import "encoding/json"

type Token string

func NewToken(in string) (out Token, err error) {
	out = Token(in)
	return
}

func (dep Token) String() string {
	return string(dep)
}

func (dep Token) MarshallJSON() ([]byte, error) {
	return json.Marshal(dep)
}
