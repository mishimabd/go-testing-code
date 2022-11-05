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

	//fmt.Println(first)
	//fmt.Println(second)
	//fmt.Println(third)
	//fmt.Println(fourth)
	//fmt.Println(fifth)
	//fmt.Println(sixth)

	switch {
	case (first + second + third) == (fourth + fifth + sixth):
		fmt.Print("YES")
	default:
		fmt.Print("NO")
	}
}
