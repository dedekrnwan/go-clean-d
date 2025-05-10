package types

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode"
)

type Phone string

func (dep Phone) MarshallJSON() ([]byte, error) {
	return json.Marshal(dep)
}

func NewPhone(number string) (Phone, error) {
	// Remove any non-digit characters
	cleaned := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, number)

	// Validate the length (assuming a 10-digit phone number for this example)
	if len(cleaned) < 10 {
		return "", fmt.Errorf("invalid phone number: must be at least 10 digits")
	}

	// If the number starts with '0', replace it with '+62'
	if strings.HasPrefix(cleaned, "0") {
		cleaned = "62" + cleaned[1:]
	}

	// If the number doesn't start with '62', prepend '+62'
	if !strings.HasPrefix(cleaned, "62") {
		cleaned = "62" + cleaned
	}

	return Phone(cleaned), nil
}

// Format formats the phone number as (XXX) XXX-XXXX
func (p Phone) Format() string {
	if len(p) < 10 {
		return string(p) // Return as-is if invalid
	}
	return fmt.Sprintf("+%s %s %s %s", p[:2], p[2:5], p[5:9], p[9:])
}

// IsValid checks if the phone number is valid
func (p Phone) IsValid() bool {
	return len(p) >= 10 && strings.HasPrefix(string(p), "62")
}

func (dep Phone) String() string {
	return string(dep)
}
