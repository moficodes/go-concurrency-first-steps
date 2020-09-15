package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	procs := runtime.GOMAXPROCS(2)
	fmt.Println(procs)

	var wg sync.WaitGroup
	stop := make(chan int)
	wg.Add(2)
	go func() {
		<-stop
		defer wg.Done()
		for number := 1; number < 1000; number++ {
			fmt.Printf("A%d\t", number)
		}
	}()

	go func() {
		<-stop
		defer wg.Done()
		for number := 1; number < 1000; number++ {
			fmt.Printf("B%d\t", number)
		}
	}()

	close(stop)
	wg.Wait()
}
