package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("moving on")
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
}
