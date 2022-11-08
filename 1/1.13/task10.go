package main

import "fmt"

func main() {
	var numberOne, numberTwo, max int
	fmt.Scan(&numberOne, &numberTwo)
	if (numberOne == 0) && (numberTwo == 0) {
		fmt.Println(0)
	}
	for i := numberTwo; i < numberOne; i-- {
		if i%7 == 0 {
			max = i
			break
		}
	}
	fmt.Println(max)
	if max != 0 {
		fmt.Println(max)
	} else {
		fmt.Println("NO")
	}

}
