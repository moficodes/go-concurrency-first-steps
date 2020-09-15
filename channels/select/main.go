package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		close(c1)
	}()

	select {
	case <-c1:
		fmt.Println("path 1")
	case c2 <- 1:
		fmt.Println("path 2")
	}
}
