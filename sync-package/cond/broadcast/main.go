package main

import (
	"fmt"
	"sync"
)

func mapCheck(data map[int]interface{}, key int) bool {
	if _, ok := data[key]; !ok {
		return true
	}
	return false
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	shared := make(map[int]interface{})
	m := &sync.Mutex{}
	c := sync.NewCond(m)

	for i := 0; i<5; i++ {
		go func(i int) {
			c.L.Lock()
			for mapCheck(shared, i) {
				c.Wait()
			}
			fmt.Println(shared[i])
			c.L.Unlock()
			wg.Done()
		}(i)
	}

	c.L.Lock()
	for i := 0; i < 5; i++ {
		shared[i] = i
	}
	c.Broadcast()
	c.L.Unlock()
	wg.Wait()
}

