package __1

func fibonacci(n int) int {
	a := 1
	b := 1
	for i := 0; i < n-1; i++ {
		a, b = b, b+a
	}
	return a
}
