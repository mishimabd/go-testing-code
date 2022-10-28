package main

import (
	"fmt"
	"sync"
)

func main() {
	waitingroup := sync.WaitGroup{}
	waitingroup.Add(4)
	go func() {
		waitingroup.Done()
		fmt.Print("1 ")
	}()
	go func() {
		waitingroup.Done()
		fmt.Print("2 ")
	}()
	go func() {
		waitingroup.Done()
		fmt.Print("3 ")
	}()
	go func() {
		waitingroup.Done()
		fmt.Print("4 ")
	}()
	waitingroup.Wait()
	fmt.Println("Exit!")
}
