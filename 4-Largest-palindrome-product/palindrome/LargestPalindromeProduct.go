package largestprime

import (
	"fmt"
	"strconv"
)

// Reverse Reverts a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(int n) bool {
	strnumber := strconv.Itoa(n)
	return strnumber == Reverse(strnumber)
}

// LargestPalindromeProduct return the largest palindrome Product
func LargestPalindromeProduct(a, b int) int {

	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			product := i * j
			fmt.Println(product)
			if isPalindrome(product) {

			}
		}
	}

	return 30
}
