package problems

func Fibonacci(n int) int {
	if n <= 2 && n >= 0 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
