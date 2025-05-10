package validator

func IsNegativeValues(values []int) (isValid bool) {
	for _, value := range values {
		if value < 0 {
			return true
		}
	}

	return false
}
