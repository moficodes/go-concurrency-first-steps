package main

import (
	"fmt"
	"time"
)

// go run -race main.go
func main() {
	data := 0

	go func() {
		data++
	}()

	if data == 0 {
		time.Sleep(1 * time.Nanosecond)
		fmt.Printf("%d\n", data)
	}
}