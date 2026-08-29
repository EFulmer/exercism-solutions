package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    if n < 1 {
        return 0, errors.New("Passed in n ≥ 1")
    }
    steps := 0
    for {
        if n == 1 {
            break
        } else {
            n = step(n)
            steps += 1
        }
    }
    return steps, nil
}

func step(n int) int {
    if n % 2 == 0 {
        return n / 2
    } else {
        return n * 3 + 1
    }
}