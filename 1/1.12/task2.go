package main

import "fmt"

func main() {
	var number int
	var array [10]int
	fmt.Scan(&number)
	for i := 0; i < number; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Println(array[3])
}
