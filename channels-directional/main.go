package main

import "fmt"

func main() {
	c := make(chan int, 2)
	// send-only
	cs := make(chan<- int, 2)
	// receive-only
	cr := make(<-chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Printf("c: %T\n", c)
	fmt.Printf("cs: %T\n", cs)
	fmt.Printf("cr: %T\n", cr)
}
