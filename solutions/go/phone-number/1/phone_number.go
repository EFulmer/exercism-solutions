package phonenumber

import (
    "errors"
    "strings"
)

func Number(phoneNumber string) (string, error) {
	filtered := filterNonNumerals(phoneNumber)
    if len(filtered) == 11 {
        return processWithCountryCode(filtered)
    } else if len(filtered) == 10 {
        return processWithoutCountryCode(filtered)
    }
    return "", errors.New("invalid phone number")
}

func filterNonNumerals(s string) string {
    var b strings.Builder
    for _, c := range s {
        asciiVal := int(c - '0')
        if asciiVal >= 0 && asciiVal <= 9 {
            b.WriteRune(c)
        }
    }
    return b.String()
}

func processWithCountryCode(phoneNumber string) (string, error) {
    // First check that the country code is 1, since that's the only one supported as described in the problem spec:
    if phoneNumber[0] != '1' {
        return "", errors.New("country code must be 1")
    } else if phoneNumber[1] == '0' || phoneNumber[1] == '1' { // area code
        return "", errors.New("area code must not start with 1")
    } else if phoneNumber[4] == '0' || phoneNumber[4] == '1' { // exchange
        return "", errors.New("exchange code must not start with 1")
    }
    return phoneNumber[1:], nil
}

func processWithoutCountryCode(phoneNumber string) (string, error) {
    if phoneNumber[0] == '0' || phoneNumber[0] == '1' { // area code
        return "", errors.New("area code must not start with 1")
    } else if phoneNumber[3] == '0' || phoneNumber[3] == '1' { // exchange code
        return "", errors.New("exchange code must not start with 1")
    }
    return phoneNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	parsedNumber, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return parsedNumber[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	parsedNumber, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    var b strings.Builder
    b.WriteRune('(')
    b.WriteString(parsedNumber[0:3])
    b.WriteString(") ")
    b.WriteString(parsedNumber[3:6])
    b.WriteRune('-')
    b.WriteString(parsedNumber[6:])
    return b.String(), nil
}
