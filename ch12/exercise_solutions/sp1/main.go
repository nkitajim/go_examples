package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func ChanSum(ch chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}

	return sum
}

func Sum(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func main() {
	bufSize := 1
	if len(os.Args) == 2 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "usage: %s <int>", os.Args[0])
		}
		bufSize = i
	}
	ch := make(chan int, bufSize)
	var wg sync.WaitGroup
	var start, end time.Time

	values := make([]int, 0, 10000000)
	for i := 0; i < 10000000; i++ {
		values = append(values, i)
	}

	start = time.Now()
	fmt.Println(Sum(values))
	end = time.Now()
	fmt.Println(end.Sub(start))

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, v := range values {
			ch <- v
		}
	}()
	go func() {
		wg.Wait()
		close(ch)
	}()

	start = time.Now()
	fmt.Println(ChanSum(ch))
	end = time.Now()
	fmt.Println(end.Sub(start))
}
