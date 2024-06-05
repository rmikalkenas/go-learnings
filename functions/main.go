package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("Foo")
}

func bar() {
	fmt.Println("Bar")
}

// func main() {
// 	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// 	fmt.Println("SUM:", sum(xi...))
// }

// func sum(ii ...int) int {
// 	fmt.Println(ii)
// 	fmt.Printf("%T\n", ii)

// 	sum := 0
// 	for _, v := range ii {
// 		sum += v
// 	}

// 	return sum
// }
