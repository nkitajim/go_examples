package main

import (
	"fmt"

	"math"
	"math/rand"
	"time"
)

func main() {
	is := make([]int, 0, 100)
	rand.Seed(time.Now().UnixNano())
	for ix := 0; ix < 100; ix++ {
		is = append(is, rand.Intn(math.MaxInt32))
	}

	fmt.Println(is)

	for _, ix := range is {
		switch {
		case 0 == ix%2 && 0 == ix%3:
			fmt.Println("Six!", ix)
		case 0 == ix%2:
			fmt.Println("Two!", ix)
		case 0 == ix%3:
			fmt.Println("Three!", ix)
		default:
			fmt.Println("Never mind!", ix)
		}
	}
}
