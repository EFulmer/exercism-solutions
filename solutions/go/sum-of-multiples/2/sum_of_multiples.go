package sumofmultiples

type Bag map[int]struct{}

func SumMultiples(limit int, divisors ...int) int {
	result := 0
	bag := Bag{}

	for _, divisor := range divisors {
		if divisor == 0 {
			continue
		}
		// Wouldn't be surprised if there's a closed form of this or at least something that makes it better than O(n^2)...
		for i := divisor; i < limit; i += divisor {
			bag[i] = struct{}{}
		}
	}
	for n := range bag {
		result += n
	}
	return result
}
