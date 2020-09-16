package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := time.Tick(1 * time.Second)
	c2 := time.Tick(2 * time.Second)
	c3 := time.Tick(3 * time.Second)

	c4 := make(chan int)
	c5 := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		close(c4)
	}()

	go func() {
		time.Sleep(3 * time.Second)
		c5 <- 42
	}()

	for {
		select {
		case <-c1:
			fmt.Println("1 Second")
		case <-c2:
			fmt.Println("2 Second")
		case <-c3:
			fmt.Println("3 Second")
		case <-c4:
			fmt.Println("Work Done")
			return
		case d := <-c5:
			fmt.Println("Data is ", d)
		}
	}
}
