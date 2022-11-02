package main

import (
	"fmt"
)

func main() {
	var firstNumber, secondNumber int
	var one, two, three, four int
	var oneS, twoS, threeS, fourS int
	fmt.Scan(&firstNumber, &secondNumber)

	//switch for first number
	switch {
	case (firstNumber > 0) && (firstNumber < 9):
		firstNumber = one
	case (firstNumber > 10) && (firstNumber < 99):
		one = firstNumber / 10
		two = firstNumber % 10
	case (firstNumber > 100) && (firstNumber < 999):
		one = firstNumber / 100
		two = firstNumber / 10 % 10
		three = firstNumber % 10
	case (firstNumber > 1000) && (firstNumber < 9999):
		one = firstNumber / 1000
		two = firstNumber / 100 % 10
		three = firstNumber / 10 % 10
		four = firstNumber % 10
	}

	//Switch for second number
	switch {
	case (secondNumber > 0) && (secondNumber < 9):
		oneS = secondNumber
	case (secondNumber > 10) && (secondNumber < 99):
		oneS = secondNumber / 10
		twoS = secondNumber % 10
	case (secondNumber > 100) && (secondNumber < 999):
		oneS = secondNumber / 100
		twoS = secondNumber / 10 % 10
		threeS = secondNumber % 10
	case (secondNumber > 1000) && (secondNumber < 9999):
		oneS = secondNumber / 1000
		twoS = secondNumber / 100 % 10
		threeS = secondNumber / 10 % 10
		fourS = secondNumber % 10
	}

	if one == oneS {
		fmt.Print(one)
	}
	if two == twoS {
		fmt.Print(two)
	}
	if three == threeS {
		fmt.Print(three)
	}
	if four == fourS {
		fmt.Print(four)
	}
}
