package main

import (
	"fmt"
	"time"
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

func onedelay() {
	time.Sleep(time.Second * 3)
	fmt.Println("I am nigger 3")
}
func second() {
	time.Sleep(time.Second * 4)
	fmt.Println("I am nigger 4")
}
func third() {
	time.Sleep(time.Second * 5)
	fmt.Println("I am nigger 5")
}
func main() {
	fmt.Println("go on")
	go onedelay()
	go second()
	go third()
	time.Sleep(time.Second * 6)
}

//func Anything(anything interface{}) {
//	fmt.Println(anything)
//}
