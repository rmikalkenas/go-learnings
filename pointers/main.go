package main

import "fmt"

func main() {
	// x := 42
	// fmt.Println(x)
	// fmt.Println(&x)
	// fmt.Printf("%T\n", x)

	// y := &x
	// fmt.Printf("%v\t%T\n", y, y)
	// fmt.Println(y)

	// *y = 43
	// fmt.Println(x)
	// fmt.Println(*y)

	a := 42
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)

	xi := []int{1, 2, 3, 4}
	fmt.Println(xi)
	sliceDelta(xi)
	fmt.Println(xi)

	m := make(map[string]int)
	m["James"] = 42
	fmt.Println(m)
	mapDelta(m)
	fmt.Println(m)

	fmt.Println("Value semantics")
	b := 1
	fmt.Println(b)         // 1
	fmt.Println(addOne(b)) // 2
	fmt.Println(b)         // 1

	fmt.Println("Pointer semantics")
	c := 1
	fmt.Println(c) // 1
	addOneP(&c)
	fmt.Println(c) // 2
}

func intDelta(n *int) {
	*n = 43
}

func sliceDelta(ii []int) {
	ii[0] = 99
}

func mapDelta(mm map[string]int) {
	mm["Rokas"] = 33
}

func addOne(v int) int {
	return v + 1
}

func addOneP(v *int) {
	*v += 1
}
