package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("sheep", c)
	//// sending and receiving on a channel is blocking operation
	// for {
	// 	msg, open := <-c
	// 	fmt.Println(msg)
	// 	if !open {
	// 		break
	// 	}
	// }

	//// better way to write this
	for msg := range c {
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		// fmt.Println(i, thing)
		c <- fmt.Sprint(i) + " " + thing
		time.Sleep(time.Millisecond * 500)
	}
	// prevents deadlock from sender side
	close(c)
}
