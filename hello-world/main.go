package main

import (
	"fmt"
	"runtime"
)

// go run main.go.go
// without un commenting line 16 this code wont do anything.
func main() {
	go func() {
		fmt.Println("Hello Closure")
	}()

	go print("Hello Function")

	// time.Sleep(1 * time.Second)
	// time.Sleep(25 * time.Microsecond)

	// main.go function exits here.
	// and which exits the context
	// this methods lets us see the number of goroutines in context right now.
	// run this code with and without the sleep and see if that number is any different.
	fmt.Println(runtime.NumGoroutine())
}

func print(message string) {
	fmt.Println(message)
}
