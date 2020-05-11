package multiples

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
