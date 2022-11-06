package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	mass := make([]int, number)
	for i := 0; i < number; i++ {
		fmt.Scan(&mass[i])
	}
	for i := 0; i < number; i++ {
		if i%2 == 0 {
			fmt.Print(mass[i], " ")
		}
	}
}
