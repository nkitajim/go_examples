package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup

	val1 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	val2 := []int{20, 21, 22, 23, 24, 25, 26, 27, 28, 29}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, v := range val1 {
			ch1 <- v
		}
		close(ch1)
	}()
	go func() {
		defer wg.Done()
		for _, v := range val2 {
			ch2 <- v
		}
		close(ch2)
	}()
	go func() {
		wg.Wait()
	}()

	cnt := 2
	for cnt != 0 {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				cnt--
				break
			}
			fmt.Println("ch1", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				cnt--
				break
			}
			fmt.Println("ch2", v)
		}
	}
}
