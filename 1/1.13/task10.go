package main

import "fmt"

func main() {
	var numberOne, numberTwo, max int
	fmt.Scan(&numberOne, &numberTwo)
	for i := numberOne; i < numberTwo; i++ {
		if i%7 == 0 {
			max = i
		}
	}
	fmt.Println(max)
}
