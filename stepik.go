package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a)
	b = a / 30
	c = (a % 30) / 2
	fmt.Print("It is ", b, " hours ", c, " minutes.")
}
