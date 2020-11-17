package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(100*time.Millisecond))

	defer cancel()
	c := make(chan int)
	go func() {
		time.Sleep(200 * time.Millisecond)
		close(c)
	}()

	select {
	case <-c:
		fmt.Println("Work Done")
	case <-ctx.Done():
		fmt.Println("deadline hit")
	}
}
