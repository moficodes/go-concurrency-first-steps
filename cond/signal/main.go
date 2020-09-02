package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]int, 0)

	serveCustomer := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		custID := queue[0]
		queue = queue[1:]

		fmt.Println("serving customer ", custID)
		c.L.Unlock()
		c.Signal()
	}

	for i := 0 ; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 5 {
			c.Wait()
		}
		time.Sleep(500 * time.Millisecond)
		fmt.Println("adding customer ", i)
		queue = append(queue, i)
		go serveCustomer(5000 * time.Millisecond)
		c.L.Unlock()
	}
}