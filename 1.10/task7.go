package main

import "fmt"

func main() {
	var StartMoney, Percent, FinishMoney, year int
	fmt.Scan(&StartMoney, &Percent, &FinishMoney)
	for {
		if StartMoney > FinishMoney {
			break
		} else {
			StartMoney = StartMoney + StartMoney*(Percent/100)
			year++
		}
	}
	fmt.Println(year)
}
