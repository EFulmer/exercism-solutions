package luhn

import "strings"

func Valid(id string) bool {
    id = strings.ReplaceAll(id, " ", "")
    if len(id) < 2 {
        return false
    }
	digits := []int{}
    for _, char := range id {
        dig := int(char - '0')
        if dig < 0 || dig > 9 {
            return false
        }
        digits = append(digits, dig)
    }
    double := false
    for i := len(digits)-1; i >= 0; i-- {
        if double {
            newDigit := digits[i] * 2
            if newDigit >= 10 {
                newDigit -= 9
            }
            digits[i] = newDigit
        }
        double = !double
    }
    sum := 0
    for _, dig := range digits {
        sum += dig
    }
    return sum % 10 == 0
}
