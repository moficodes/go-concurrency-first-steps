package main

import (
	"fmt"
	"runtime"
	"sync"
)

// go run main.go.go

var wg sync.WaitGroup

/*
type WaitGroup struct {
	noCopy noCopy

	// 64-bit value: high 32 bits are counter, low 32 bits are waiter count.
	// 64-bit atomic operations require 64-bit alignment, but 32-bit
	// compilers do not ensure it. So we allocate 12 bytes and then use
	// the aligned 8 bytes in them as state, and the other 4 as storage
	// for the sema.
	state1 [3]uint32
}
 */

func main() {
	// fmt.Printf("%#v\n", wg)
	// try changing this number to something other than 2
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
