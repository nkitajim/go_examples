package main

import (
	"fmt"
)

type ValidTypes interface {
	// ~int | ~float64
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func Double[T ValidTypes](i T) T {
	return i * 2
}

func main() {
	fmt.Println(Double(2))
	fmt.Println(Double(2.1))
}
