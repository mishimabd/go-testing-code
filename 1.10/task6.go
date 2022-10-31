package main

import (
	"fmt"
)

func main() {
	var b int
	for fmt.Scan(&b); b <= 100; fmt.Scan(&b) {
		if (b >= 10) && (b <= 100) {
			fmt.Println(b)
			continue
		}

	}
}
