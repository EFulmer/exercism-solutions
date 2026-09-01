package darts

import "math"

func Score(x, y float64) int {
	distance := math.Sqrt(x*x + y*y)
    switch {
    case distance <= 1:
        return 10
    case 1 < distance && distance <= 5:
        return 5
    case 5 < distance && distance <= 10:
        return 1
    default:
        return 0
    }
}
