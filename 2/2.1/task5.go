package __1

func sumInt(args ...int) (int, int) {
	sum := 0
	for _, x := range args {
		sum += x
	}
	return len(args), sum
}
