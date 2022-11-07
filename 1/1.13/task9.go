package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	fmt.Println((number-1)%9 + 1)
}
