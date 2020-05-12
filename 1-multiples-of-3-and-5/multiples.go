package multiples

// 1) Multiples of 3 and 5
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.
//FindMultiples sum the multiples of 3 and 5 until n
func FindMultiples(n int) int {
	out := 0
	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			out += i
		}
	}
	return out
}