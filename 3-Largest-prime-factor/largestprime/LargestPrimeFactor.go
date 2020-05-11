package largestprime

func isPrime(n int) bool {
	for j := 2; j < n; j++ {
		if n%j == 0 {
			return false
		}
	}
	return true
}

// LargestPrimeFactor return the largest prime factor until n
func LargestPrimeFactor(n int) int {
	var out int
	for i := 1; i*i < n; i++ {
		if n%i == 0 {
			if isPrime(i) {
				out = i
			}
		}
	}
	return out
}
