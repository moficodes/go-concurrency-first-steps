package main

import "fmt"

func main() {
	c := make(chan int, 1)

	c <- 1

	d := <-c

	fmt.Println(d)
}
