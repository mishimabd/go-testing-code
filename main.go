package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "my name is"
	b := "is"
	fmt.Println(strings.Contains(a, b))
}
