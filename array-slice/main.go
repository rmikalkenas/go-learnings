package main

import "fmt"

func main() {
	aa := []string{"James", "Bond", "Shaken, no retired"}
	bb := []string{"Miss", "Moneypenny", "I'm 008."}

	xi := [][]string{aa, bb}

	for i, inner := range xi {
		fmt.Printf("Key: %v\n", i)

		for _, value := range inner {
			fmt.Println(value)
		}
	}
}
