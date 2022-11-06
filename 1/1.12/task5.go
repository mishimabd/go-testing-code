package main

import "fmt"

func main() {
	var number, counter int
	fmt.Scan(&number)
	mass := make([]int, number)
	for i := 0; i < number; i++ {
		fmt.Scan(&mass[i])
	}
	for i := 0; i < number; i++ {
		if mass[i] > 0 {
			counter++
		}
	}
	fmt.Println(counter)
}
