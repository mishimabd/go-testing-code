package main

import "fmt"

func main() {
	var number, sum int
	fmt.Scan(&number)
	sum = number%10 + number/10%10 + number/100
	fmt.Println(sum)
}
