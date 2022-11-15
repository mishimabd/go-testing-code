package __1

import "fmt"

func minimumFromFour() int {
	var a, b, c, d, min int
	fmt.Scan(&a, &b, &c, &d)
	if (a <= b) && (a <= c) && (a <= d) {
		min = a
	}
	if (b <= a) && (b <= c) && (b <= d) {
		min = b
	}
	if (c <= b) && (c <= a) && (c <= d) {
		min = c
	}
	if (d <= b) && (d <= c) && (d <= a) {
		min = d
	}
	return min
}
