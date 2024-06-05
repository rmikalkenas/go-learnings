package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Anonymous func ran")
	}()

	call := func(s string) string {
		return fmt.Sprint("From anonoymous func: ", s)
	}

	fmt.Println(callFunction(call))

	// inc1 := incrementor()
	// inc2 := incrementor()
	// fmt.Println(inc1())
	// fmt.Println(inc1())
	// fmt.Println(inc2())
	// fmt.Println(inc1())
	// fmt.Println(inc2())

	fmt.Println(factorial(4))
}

func factorial(n int) int {
	if n <= 1 {
		return n
	}

	return n * factorial(n-1)
}

func foo() {
	fmt.Println("Foo ran")
}

func callFunction(c func(s string) string) string {
	return c("HI!!!")
}

func incrementor() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}
