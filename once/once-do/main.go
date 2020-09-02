package main

import (
	"fmt"
	"sync"
)

// uncomment the wg line to see it run as go routines
func main() {
	var once sync.Once
	//var wg sync.WaitGroup

	onceTask := func(i int) func() {
		return func() {
			fmt.Println(i)
		}
	}

	//wg.Add(100)
	for i := 0; i < 100; i++ {
		//go func() {
		//defer wg.Done()
		once.Do(onceTask(i))
		//}()
	}
	//wg.Wait()
}
