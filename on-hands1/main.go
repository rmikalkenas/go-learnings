package main

import (
	"fmt"
)

func main() {
	var a int8 = 1<<7 - 1
	var b uint8 = 1<<8 - 1

	fmt.Println("😁")
	fmt.Printf("%d %d", a, b)
}
