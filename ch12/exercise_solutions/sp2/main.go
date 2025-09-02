package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type SquareMap map[int]float64

func buildSquareRootMap() SquareMap {
	out := make(SquareMap, 100_000)
	for i := 0; i < 100_000; i++ {
		out[i] = math.Sqrt(float64(i))
	}
	return out
}

var squareRootMapCache = sync.OnceValue(buildSquareRootMap)

func main() {
	var sum float64
	var start, end time.Time

	sm := buildSquareRootMap()

	start = time.Now()
	sum = 0
	for i := 0; i < 100000; i++ {
		sum += sm[i]
	}
	end = time.Now()
	fmt.Println("map", sum, end.Sub(start))

	start = time.Now()
	sum = 0
	for i := 0; i < 100000; i++ {
		sum += squareRootMapCache()[i]
	}
	end = time.Now()
	fmt.Println("cache1", sum, end.Sub(start))

	start = time.Now()
	sum = 0
	for i := 0; i < 100000; i++ {
		sum += squareRootMapCache()[i]
	}
	end = time.Now()
	fmt.Println("cache2", sum, end.Sub(start))
}
