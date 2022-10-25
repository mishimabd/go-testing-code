package main

import "fmt"

type Nigger struct {
	price int
	name  string
}

type Acts interface {
	Working()
	NotEating()
	EatChicken()
}

func (n *Nigger) Working(act string) {
	fmt.Println("I'm", act, "damn")
}

//func (n *Nigger) getName() {
//	fmt.Println(n.name)
//}
//
//func (an *Nigger) getPrice() {
//	fmt.Println(an.price)
//}

func main() {

	n := Nigger{
		name:  "Kaley",
		price: 124,
	}
	an := Nigger{
		price: 500,
		name:  "Zoki",
	}
	var a, b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	n.Working(a)
	an.Working(b)
}
