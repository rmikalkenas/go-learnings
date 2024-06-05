package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(eve, odd, quit)

	// receive
	receive(eve, odd, quit)

	fmt.Println("about to exit")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	close(q)
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel:", v)
		case v := <-o:
			fmt.Println("from the odd channel:", v)
		case v, ok := <-q:
			if !ok {
				fmt.Println("from the quit channel !ok:", v, ok)
			} else {
				fmt.Println("from the quit channel:", v, ok)
			}
			return
		}
	}
}
