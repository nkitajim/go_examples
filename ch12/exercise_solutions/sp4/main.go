package main

import (
	"context"
	// "fmt"
)

// this code goroutine leak
func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func countToEx(ctx context.Context, max int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < max; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()
	return ch
}

func main() {
	for {
		func() {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			for i := range countToEx(ctx, 10) {
				if i > 5 {
					break
				}
			}
		}()
	}

	/* leak code
	for {
		for i := range countTo(10) {
			if i > 5 {
				break
			}
		}
	}
	*/
}
