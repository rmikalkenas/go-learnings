package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}

func (d *dog) run() {
	d.first = "Rover"
	fmt.Println("My name is", d.first, "and I'm running.")
}

type yougin interface {
	walk()
	run()
}

func yougRun(y yougin) {
	y.walk()
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	// d1.walk()
	// d1.run()
	// fmt.Println(d1)
	yougRun(&d1)

	d2 := &dog{"Padget"}
	// d2.walk()
	// d2.run()
	// fmt.Println(d2)
	yougRun(d2)
}
