package main

import (
	"fmt"
	"math"
	"sync"
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
	for i := 0; i < 100000; i += 1000 {
		fmt.Println(i, squareRootMapCache()[i])
	}
}
