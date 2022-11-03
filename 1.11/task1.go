package main

import "fmt"

func main() {
	var a, num float64
	fmt.Scan(&a)
	num = a * a
	if a <= 0 {
		fmt.Printf("число %.2f не подходит", a)
	} else if a > 9999 {
		fmt.Printf("%e", a)
	} else {
		fmt.Printf("%1.4f", num)
	}
}
