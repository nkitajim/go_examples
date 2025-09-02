package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func Sum(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func GoMutexSum(values []int, result *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	sum := 0
	for _, v := range values {
		sum += v
	}

	mu.Lock()
	*result += sum
	mu.Unlock()
}

func GoAtomicSum(values []int, result *atomic.Int32, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	for _, v := range values {
		sum += v
	}

	result.Add(int32(sum))
}

func main() {
	var start, end time.Time

	// create input data
	values := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		values = append(values, i)
	}

	// normal
	{
		start = time.Now()
		sum := 0
		for i := 0; i < 10; i++ {
			sum += Sum(values)
		}
		end = time.Now()
		fmt.Println("normal", sum, end.Sub(start))
	}

	// go routine mutex
	{
		var wg sync.WaitGroup
		var mu sync.Mutex

		start = time.Now()
		sum := 0
		wg.Add(10)
		for i := 0; i < 10; i++ {
			go func() { GoMutexSum(values, &sum, &wg, &mu) }()
		}
		wg.Wait()

		end = time.Now()
		fmt.Println("go routine mutex", sum, end.Sub(start))
	}

	// go routine atomic
	{
		var wg sync.WaitGroup

		start = time.Now()
		var sum atomic.Int32
		wg.Add(10)
		for i := 0; i < 10; i++ {
			go func() { GoAtomicSum(values, &sum, &wg) }()
		}
		wg.Wait()

		end = time.Now()
		fmt.Println("go routine atomic", sum.Load(), end.Sub(start))
	}

	// go routine channel
	{
		ch := make(chan int)
		var wg sync.WaitGroup

		start = time.Now()
		sum := 0
		wg.Add(10)
		for i := 0; i < 10; i++ {
			go func() {
				defer wg.Done()
				ch <- Sum(values)
			}()
		}
		go func() {
			wg.Wait()
			close(ch)
		}()
		for v := range ch {
			sum += v
		}
		end = time.Now()
		fmt.Println("go routine channel", sum, end.Sub(start))
	}
}
