package __1

func vote(x int, y int, z int) int {
	var answer int
	if x == y {
		answer = x
	}
	if x == z {
		answer = x
	}
	if y == z {
		answer = y
	}
	return answer
}
