package main

import "fmt"

func main() {
	var workArray [10]uint8
	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&workArray[i])
	}
	for i := 0; i < 3; i++ {
		var first, second int
		fmt.Scan(&first, &second)
		numberOne := workArray[first]
		secondNumber := workArray[second]
		workArray[first] = secondNumber
		workArray[second] = numberOne
	}
	for i := 0; i < len(workArray); i++ {
		fmt.Print(workArray[i], " ")
	}

}
