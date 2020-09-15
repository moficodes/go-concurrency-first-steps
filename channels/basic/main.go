package main

import "fmt"

func main() {
	c := make(chan int)

	c <- 1

	d := <-c

	fmt.Println(d)
}

//go func() {
//	c <- 1
//}()
