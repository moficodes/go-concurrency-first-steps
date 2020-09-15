package main

import "fmt"

func main() {
	c := make(chan int)

	go send(c)

	d := receive(c)

	fmt.Println(d)
}

func send(c chan<- int) {
	//d := <-c

	c <- 1
}

func receive(c <-chan int) int {
	//c <- 1
	//close(c)
	return <-c
}
