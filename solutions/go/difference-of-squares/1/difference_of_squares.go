package differenceofsquares

func SquareOfSum(n int) int {
    sumN := n * (n + 1) / 2
    return sumN * sumN
}

func SumOfSquares(n int) int {
    result := 0
    for i := range n+1 {
        result += (i * i)
    }
    return result
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
