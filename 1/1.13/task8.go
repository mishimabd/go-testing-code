package main

import "fmt"

func main() {
	var number, counter, min, temp int
	fmt.Scan(&number)
	for i := 0; i < number; i++ {
		fmt.Scan(&temp)
		if min > temp {
			min = temp
			counter = 1
		}
		if min == temp {
			counter++
		}
	}
	fmt.Println(counter)
}
