package main

import "fmt"

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// concurrent goroutine
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	// fill up jobs channel
	for i := 0; i < 100; i++ {
		jobs <- i
	}
	// close jobs channel
	close(jobs)
	// expect 100 results
	for j := 0; j < 100; j++ {
		fmt.Print(j, " ")
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
