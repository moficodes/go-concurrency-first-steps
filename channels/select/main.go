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
		//default:
		//	fmt.Println("nothing else was happening")
	}

	//count1, count2, count3 := 0, 0, 0
	//close(c2)
	//for i := 0; i < 10000; i++ {
	//	select {
	//	case <-c1:
	//		count1++
	//	case <-c2:
	//		count2++
	//	default:
	//		count3++
	//	}
	//}
	//
	//fmt.Printf("count1: %d, count2: %d, count3: %d\n", count1, count2, count3)
}
