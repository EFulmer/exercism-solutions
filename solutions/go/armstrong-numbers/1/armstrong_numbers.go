package armstrongnumbers

import "strconv"

func IsNumber(n int) bool {
	digits := len(strconv.Itoa(n))
	sum := 0
	for _, c := range strconv.Itoa(n) {
		curDigit := int(c - '0')
		sum += pow(curDigit, digits)
	}
	return sum == n
}

func pow(m, n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= m
	}
	return result
}
