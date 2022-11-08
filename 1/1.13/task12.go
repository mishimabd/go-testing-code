package main

import "fmt"

func main() {
	var number int
	two := 2
	temp := two * 2
	fmt.Scan(&number)
	if number == 2 || number == 1 {
		fmt.Println(1)
	} else {
		fmt.Print("1 2 ")
		for i := 0; i < number; i++ {
			if temp > number {
				break
			} else {
				fmt.Print(temp, " ")
				temp = temp * 2
			}
		}
	}

}
