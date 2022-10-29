package main

import "fmt"

func main() {
	var a, c int
	fmt.Scan(&a)
	for i := 1; i <= a; i++ {
		fmt.Scan(&i)
		if (i%8 == 0) && ((i > 10) && (i < 99)) {
			fmt.Println("Successfully")
			c += i
		}
		fmt.Println("Successfully two")
	}
	fmt.Println("Successfully three")
	fmt.Println(c)
}
