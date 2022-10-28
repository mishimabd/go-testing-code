package main

import (
	"fmt"
)

//type Nigger struct {
//	price int
//	name  string
//}

//
//func (n *Nigger) NotEating() {
//	fmt.Println("I am not eating")
//}
//
//func (n *Nigger) EatChicken() {
//	fmt.Println("I am eating chicken!")
//}
//
//type Acts interface {
//	Working()
//	NotEating()
//	EatChicken()
//}
//
//func NewModel(b int, arg string) *Acts {
//	return &Nigger{}
//}
//
//func (n *Nigger) Working() {
//	fmt.Println("I'm damn")
//}

//func (n *Nigger) getName() {
//	fmt.Println(n.name)
//}
//
//func (an *Nigger) getPrice() {
//	fmt.Println(an.price)
//}

//func onedelay() {
//	time.Sleep(time.Second * 3)
//	fmt.Println("I am nigger 3")
//}
//func second() {
//	time.Sleep(time.Second * 4)
//	fmt.Println("I am nigger 4")
//}
//func third() {
//	time.Sleep(time.Second * 5)
//	fmt.Println("I am nigger 5")
//}
//func left() {
//	for {
//		fmt.Println("I am nigger which is leftover")
//		time.Sleep(time.Second * 1)
//	}
//}

func main() {
	fmt.Print("Hello ")
	go func() {
		fmt.Print("World ")
	}()
	fmt.Println("!")
}

//func Anything(anything interface{}) {
//	fmt.Println(anything)
//}
