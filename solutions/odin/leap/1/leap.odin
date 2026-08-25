package leap

is_leap_year :: proc(year: int) -> bool {
    switch {
    case divides(year, 100) && !divides(year, 400):
        return false
    case divides(year, 4):
        return true
    case:
        return false
    }
}

divides :: proc(m, n: int) -> bool {
    return m % n == 0
}
