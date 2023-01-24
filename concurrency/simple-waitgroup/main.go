package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//// these will run side-by-side
	// go count("cat")
	// count("fish")

	//// using waitgroups
	var wg sync.WaitGroup
	wg.Add(1) // we're going to wait for one goroutine

	go func() {
		count("sheep")
		wg.Done() // we're done waiting for this goroutine. This decrements counter by 1
	}()

	// this blocks until the counter is 0
	wg.Wait()
}

func count(thing string) {
	for i := 0; i <= 10; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
