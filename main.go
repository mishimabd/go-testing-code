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

func heavy() {
	time.Sleep(time.Second * 5)
}
func main() {
	go heavy()
	fmt.Println("Nigger")

}

//func Anything(anything interface{}) {
//	fmt.Println(anything)
//}
