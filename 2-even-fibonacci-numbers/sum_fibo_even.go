package fibo

// SumFiboEven sum all the even number in fibonaccy sequence until n
func SumFiboEven(n int) int {
	a, b := 1, 1
	sum := 0

	for b <= n {
		a, b = b, a+b
		if b%2 == 0 {
			sum += b
		}
	}

	return sum
}
