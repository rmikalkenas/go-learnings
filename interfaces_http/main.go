package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(os.Stdout, resp.Body)
}

func (logWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	fmt.Println("JUST WROTE THIS MANY BYTES:", len(p))

	return len(p), nil
}
