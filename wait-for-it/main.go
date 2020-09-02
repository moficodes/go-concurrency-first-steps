package main

import (
	"fmt"
	"runtime"
	"sync"
)

// go run main.go

var wg sync.WaitGroup

func main() {
	// fmt.Printf("%#v\n", wg)
	wg.Add(2)
	// fmt.Printf("%#v\n", wg)
	go func() {
		// defer wg.Done()
		fmt.Println("Hello Closure")
		wg.Done()
	}()

	go print("Hello Function")

	// fmt.Printf("%#v\n", wg)
	wg.Wait()
	// fmt.Printf("%#v\n", wg)
	fmt.Println(runtime.NumGoroutine())
}

func print(message string) {
	// defer wg.Done()
	fmt.Println(message)
	wg.Done()
}
