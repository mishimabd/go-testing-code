package main

import "fmt"

type Nigger struct {
	price int
	name  string
}

func (n *Nigger) getName() {
	fmt.Println(n.name)
}

func (an *Nigger) getPrice() {
	fmt.Println(an.price)
}

func main() {
	n := Nigger{
		name:  "Kaley",
		price: 124,
	}
	an := Nigger{
		price: 500,
		name:  "Zoki",
	}
	fmt.Println(an, n)
}
