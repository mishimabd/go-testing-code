package main

import "fmt"

func main() {
	var time, hours, minutes int
	fmt.Scan(&time)
	hours = time / 3600
	minutes = time % 3600 / 60
	fmt.Printf("It is %d hours %d minutes.", hours, minutes)

}
