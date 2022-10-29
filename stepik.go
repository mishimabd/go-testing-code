package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)
	first := number / 100000
	second := number / 10000 % 10
	third := number / 1000 % 10
	fourth := number / 100 % 10
	fifth := number / 10 % 10
	sixth := number % 10
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(fourth)
	fmt.Println(fifth)
	fmt.Println(sixth)
	switch {
	case (number > 0) && (number < 9):
		fmt.Print(number)
	case (number > 10) && (number < 99):
		fmt.Print(number / 10)
	case (number > 100) && (number < 999):
		fmt.Print(number / 100)
	case (number > 1000) && (number < 10000):
		fmt.Print(number / 1000)
	default:
		fmt.Print(1)
	}
}
