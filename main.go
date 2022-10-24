package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(array); i++ {
		if (array[i] % 2) == 0 {
			fmt.Print(array[i])
			fmt.Println(" It's even number")
		}

	}
}
