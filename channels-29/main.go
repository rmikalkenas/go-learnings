package main

import "fmt"

func main() {
	runBufferBlock()
}

func runBufferBlock() {
	c := make(chan int, 1)

	c <- 42
	c <- 43

	fmt.Println(<-c)
}

func runSuccessfulBuffer() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}

func runSuccessful() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func doesNotRun() {
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}
