package main

import "fmt"

func main() {
	// default is unbuffered channel
	// create a buffered channel with capacity of 3 for asynchronous communication
	// this uses the queue like structure
	charChannel := make(chan string, 3)

	chars := []string{"a", "b", "c"}
	for _, s := range chars {
		select {
		case charChannel <- s:
		}
	}
	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}
}
