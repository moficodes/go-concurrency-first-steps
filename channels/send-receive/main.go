package main

import "fmt"

func main() {
	c := make(chan int)

	go send(c)

	d := receive(c)

	fmt.Println(d)
}

func send(c chan<- int) {
	c <- 1
}

func receive(c <-chan int) int {
	return <-c
}
