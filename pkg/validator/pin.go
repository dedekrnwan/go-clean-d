package validator

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	DEFAULT_SERIES_THRESHOLD = 4
)

func IsPinNumeric(value string) (bool, error) {
	valid, err := regexp.MatchString("^[0-9]{6}$", value)
	if err != nil {
		return false, err
	}
	repeat := regexp.MustCompile(`^(\d)1+$`).MatchString(value)
	return valid && !strings.Contains("0123456789876543210", value) && !repeat, err
}

func IsPinRepetitiveAndSequence(pin string, threshold int) bool {
	if threshold == 0 {
		threshold = DEFAULT_SERIES_THRESHOLD
	}
	intPIN := make([]int, len(pin))
	for idx, value := range pin {
		temp, _ := strconv.Atoi(string(value))
		intPIN[idx] = temp
	}

	increment := []int{-2, -1, 0, 1, 2}
	count_sequence := make([]int, len(increment))
	count_temp := make([]int, len(increment))
	max_value := 0

	for idx_i, value := range intPIN {
		if idx_i == 0 {
			continue
		}
		for idx_j, incr := range increment {
			if value == (intPIN[idx_i-1] + incr) {
				count_temp[idx_j] += 1
			} else {
				count_temp[idx_j] = 0
			}
			if count_sequence[idx_j] < count_temp[idx_j] {
				count_sequence[idx_j] = count_temp[idx_j]
			}
		}
	}
	for _, value := range count_sequence {
		if max_value < value {
			max_value = value
		}
	}
	return max_value >= threshold
}
