package main

import "fmt"

func main() {
	var number, counter, san int
	fmt.Scan(&number)
	for i := 0; i < number; i++ {
		fmt.Scan(&san)
		if san == 0 {
			counter++
		}
	}
	fmt.Println(counter)
}
