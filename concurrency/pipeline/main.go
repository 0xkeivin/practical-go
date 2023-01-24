package main

import "fmt"

func main() {

	// input
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// stage 1
	dataChannel := sliceToChannel(nums)
	// stage 2
	finalChannel := sq(dataChannel)
	// stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}

	// manual
	// for _, v := range nums {
	// 	fmt.Println(v * v)
	// }
}

// returns read-only channel
func sliceToChannel(nums []int) <-chan int {
	// unbuffered channel
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
