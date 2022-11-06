package main

import "fmt"

func main() {
	var first, second float32
	var average float32
	fmt.Scan(&first, &second)
	average = (first + second) / 2
	fmt.Println(average)
}
