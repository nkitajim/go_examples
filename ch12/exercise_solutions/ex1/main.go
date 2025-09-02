package main

import (
	"fmt"
	"sync"
)

func processData() {
	ch := make(chan int)

	var wg sync.WaitGroup

	val1 := []int{1, 2, 3, 4, 5}
	val2 := []int{6, 7, 8, 9, 10}

	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("Start Goroutine1")
		for _, v := range val1 {
			ch <- v
		}
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Start Goroutine2")
		for _, v := range val2 {
			ch <- v
		}
	}()
	go func() {
		wg.Wait()
		fmt.Println("Close channel")
		close(ch)
	}()

	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()

	wg2.Wait()
}

func main() {
	processData()
}
