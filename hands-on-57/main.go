package main

import (
	"fmt"
	"time"
)

func main() {
	timeFunc(doWork)
}

func doWork() {
	for i := 0; i < 2_000; i++ {
		fmt.Print(i)
	}
	fmt.Println("-----")

	
}

func timeFunc(f func()) {
	start := time.Now()
	f()
	fmt.Println(time.Since(start))
}
