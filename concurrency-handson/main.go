package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	var counter int64 = 0

	jobsCount := 1000

	wg.Add(jobsCount)

	// var mu sync.Mutex

	for i := 0; i < jobsCount; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			// mu.Lock()
			// value := counter
			// runtime.Gosched()
			// value++
			// counter = value

			// mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
