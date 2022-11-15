package main

import "fmt"

func main() {
	var number, counter, min, temp int
	fmt.Scan(&number)
	counter = 1
	for i := 0; i < number; i++ {
		fmt.Scan(&temp)

		if min < temp {
			min = temp
		}
		if temp == min {
			counter = counter + 1
		}
	}
	fmt.Println(counter)
}
