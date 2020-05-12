package main

// # 3) Largest prime factor
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

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

func isPrime(n int) bool {
	for j := 2; j < n; j++ {
		if n%j == 0 {
			return false
		}
	}
	return true
}
