package main

import "fmt"

func main() {
	var number, first, second, third int
	fmt.Scan(&number)
	first = number / 100
	second = number / 10 % 10
	third = number % 10
	fmt.Printf("%d%d%d", third, second, first)
}
