package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "password123"

	start := time.Now()
	pb, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(pb))
	end := time.Now()
	fmt.Println(end.Sub(start))

	a := "Password123"
	err = bcrypt.CompareHashAndPassword(pb, []byte(a))

	if err != nil {
		fmt.Println("You cannot login!")
		return
	}

	fmt.Println("Welcome, you are logged in")
}
