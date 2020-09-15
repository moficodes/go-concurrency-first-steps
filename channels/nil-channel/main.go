package main

import (
	"fmt"
)

func main() {
	var c chan int
	//close(c) //closing a nil channel will block

	go func() {
		c <- 1
	}()
	d := <-c

	fmt.Println(d)
}
