package main

import (
	"fmt"
	"sync"
)

func main() {
	ex7()
}

func ex7() {
	c := make(chan int)
	var wg sync.WaitGroup

	go func() {
		for g := 0; g < 10; g++ {
			wg.Add(1)

			go func(g int) {
				from := g * 10
				for i := from; i < from+10; i++ {
					c <- i
				}

				wg.Done()
			}(g)
		}

		wg.Wait()
		close(c)
	}()

	total := 0
	for v := range c {
		fmt.Println(v)
		total++
	}

	fmt.Println("Total:", total)

}

func ex6() {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}

func ex5() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

func ex4() {
	q := make(chan int)
	c := gen4(q)

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}

	}

	fmt.Println("about to exit")
}

func gen4(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func ex3() {
	c := gen3()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func gen3() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

		close(c)
	}()

	return c
}

func ex2() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

func ex1() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	// c <- 42

	fmt.Println(<-c)
}
