package main

import (
	"fmt"
	"runtime"
	"sync"
)

// go run main.go

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("Hello Closure")
		// wg.Done()
	}()

	go print("Hello Function")

	wg.Wait()
	fmt.Println(runtime.NumGoroutine())
}

func print(message string) {
	defer wg.Done()
	fmt.Println(message)
	// wg.Done()
}
