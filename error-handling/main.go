package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func main() {
	// _, err := os.Open("no-file.txt")
	// if err != nil {
	// 	// fmt.Println("Err happened", err)
	// 	fmt.Errorf()
	// 	// log.Println("Err happened", err)
	// 	// log.Fatalln(err)
	// 	log.Panicln(err)
	// 	// panic(err)
	// }

	c1 := customErr{
		info: "need more coffee",
	}

	foo(c1)
}

func foo(e error) {
	fmt.Printf("%#v\n", e)
	fmt.Println("foo ran -", e, "\n", e.(customErr).info)
}
