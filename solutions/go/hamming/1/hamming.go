package hamming

import "errors"

func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
        return 0, errors.New("strings have different lengths")
    }
    result := 0
    for i := range(len(a)) {
        if a[i] != b[i] {
            result += 1
        }
    }
    return result, nil
}
