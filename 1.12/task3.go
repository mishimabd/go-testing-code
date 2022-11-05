package main

import "fmt"

func main() {
	var mass [5]int
	for i := 0; i < len(mass); i++ {
		fmt.Scan(&mass[i])
	}
	fmt.Println(findMax(mass))
}

func findMax(arr [5]int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
