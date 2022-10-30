package main

import (
	"fmt"
)

func main() {
	var a, b, sum int
	a = 1
	sum = 0
	for fmt.Scan(&b); b != 0; fmt.Scan(&b) {
		if a < b {
			a = b
			sum = 1
		}
		if a == b {
			sum = sum + 1
		}
	}
	fmt.Println(sum - 1)
}
