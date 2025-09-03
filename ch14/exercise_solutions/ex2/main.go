package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func check1234(ctx context.Context) (int, int, error) {
	rand.Seed(time.Now().UnixNano())
	total := 0
	count := 0
	for {
		select {
		case <-ctx.Done():
			return total, count, errors.New("Timeout")
		default:
		}

		num := rand.Intn(100_000_000)
		total += num
		count++

		if 1234 == num {
			return total, count, nil
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	start := time.Now()
	total, count, err := check1234(ctx)
	end := time.Now()
	if err != nil {
		fmt.Printf("total %d, count %d, time: %s, %s\n", total, count, end.Sub(start), err)
	} else {
		fmt.Printf("total %d, count %d, time: %s, Success!\n", total, count, end.Sub(start))
	}
}
