package main

import (
	"context"
	"fmt"
	"time"
)

func doSomethingCool(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timed out")
			return
		default:
			fmt.Println("doing something cool")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Go context tutorial")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// goroutine
	go doSomethingCool(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("oh no, i've exceeded the deadline")
		// error will show non-nil once context is Done()
		fmt.Println(ctx.Err())
	}
	time.Sleep(2 * time.Second)
}

/* Output:
Go context tutorial
doing something cool
doing something cool
doing something cool
doing something cool
oh no, i've exceeded the deadline
context deadline exceeded
timed out
*/
