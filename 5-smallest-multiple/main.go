package main

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10
// without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

func SmallestMultiple(upperBound int) int {
	for x := 1; true; x++ {
		if IsDivisibleByRange(x, upperBound) {
			return x
		}
	}
}

func IsDivisibleByRange(testd int, max int) bool {
	for y := 1; y < max; y++ {
		if testd%y != 0 {
			return false
		}
	}
	return true
}
