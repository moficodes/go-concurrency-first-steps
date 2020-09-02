package main

import "sync"

// Example from https://learning.oreilly.com/library/view/concurrency-in-go/9781491941294/
func main() {
	var onceA, onceB sync.Once
	var initB func()
	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) }
	onceA.Do(initA)
}