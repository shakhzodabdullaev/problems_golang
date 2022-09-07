package problems

func Fibonacci(n int) int {
	var (
		first  = 0
		second = 1
		result = 0
	)

	for i := 0; i < n; i++ {
		result = first + second
		first = second
		second = result
	}

	return result
}
