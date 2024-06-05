package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}

func main() {
	p1 := person{
		first: "Jenny",
	}

	p1.speak()
}
