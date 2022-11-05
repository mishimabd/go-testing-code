package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := make([]int, 3, 3)
	fmt.Println(b)
	n := copy(b, a)

	fmt.Println(n)
	fmt.Printf("a = %v\n", a) // a = [1 2 3]
	fmt.Printf("b = %v\n", b) // b = [1 2 3]
	fmt.Printf("Скопировано %d элемента\n", n)
}
