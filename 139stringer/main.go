package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("Book with title: ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is the number", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer) {
	log.Println("LOG FROM 139", s.String())
}

func main() {
	b := book{
		title: "Harry Potter",
	}

	var m count = 42

	logInfo(b)
	logInfo(m)
}
