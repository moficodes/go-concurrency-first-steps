package main

import (
	"fmt"
	"time"
)

type data struct {
	val int
	//mux sync.Mutex
}

func main() {
	d := data{val: 0}

	for i := 0; i < 10000; i++ {
		go inc(&d)
	}
	time.Sleep(1 * time.Second)

	//var wg sync.WaitGroup
	//wg.Add(10000)
	//for i := 0; i < 10000; i++ {
	//	go inc(&d, &wg)
	//}
	//wg.Wait()

	//for i := 0; i < 10000; i++ {
	//	go inc(&d)
	//}
	//time.Sleep(time.Second)
	fmt.Println(d.val)
}

func inc(data *data) {
	data.val++
}

//func inc(data *data, wg *sync.WaitGroup) {
//	defer wg.Done()
//	data.val++
//}

//func inc(data *data) {
//	data.mux.Lock()
//	data.val++
//	data.mux.Unlock()
//}
