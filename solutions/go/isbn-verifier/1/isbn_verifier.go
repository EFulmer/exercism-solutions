package isbnverifier

import (
    "log"
    "regexp"
)

func IsValidISBN(isbn string) bool {
    match, err := regexp.MatchString("^\\d-?\\d{3}-?\\d{5}-?[0-9X]$", isbn) 
    if err != nil {
        log.Fatalf("Error parsing argument %s; %v\n", isbn, err)
    }
    if !match {
        return false
    }
	acc := 0
    mult := 10
    for _, v := range isbn {
        if v == '-' {
            continue
        }
        var dv int
        // technically 'X' is only allowed at the last index, but the regex catches that so this will never produce a logical error.
        if v == 'X' {
            dv = 10
        } else {
            dv = int(v - '0')
        }
        acc = acc + (dv * mult)
        mult -= 1
    }
    return (acc % 11) == 0
}
