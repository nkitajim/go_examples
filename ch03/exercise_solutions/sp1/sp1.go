package main

import (
	"fmt"
	"time"
)

type Ssi struct {
	a int
	b int
	c int
	d int
	e int
	f int
	g int
	h int
	i int
}

// const li = []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // (value of type []int) is not constant

func main() {
	var li = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var msi = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
	}
	var ssi = Ssi{1, 2, 3, 4, 5, 6, 7, 8, 9}

	var start time.Time
	var end time.Time

	fmt.Println(li)
	fmt.Println(msi)
	fmt.Println(ssi)

	start = time.Now()
	for i := 0; i < 10000000; i++ {
		sum := 0
		for ix := 0; ix < len(li); ix++ {
			sum += ix
		}
	}
	end = time.Now()
	fmt.Println("for time:", end.Sub(start))

	start = time.Now()
	for i := 0; i < 10000000; i++ {
		sum := 0
		for _, ix := range li {
			sum += ix
		}
	}
	end = time.Now()
	fmt.Println("for list time:", end.Sub(start))

	start = time.Now()
	for i := 0; i < 10000000; i++ {
		sum := 0
		for _, ix := range msi {
			sum += ix
		}
	}
	end = time.Now()
	fmt.Println("for map time:", end.Sub(start))

	// check read time
	{
		sum := 0
		start = time.Now()
		for i := 0; i < 10000000; i++ {
			sum += 1
		}
		end = time.Now()
		fmt.Println("int read time:", end.Sub(start))
	}

	{
		sum := 0
		start = time.Now()
		for i := 0; i < 10000000; i++ {
			sum += li[0]
		}
		end = time.Now()
		fmt.Println("list read time:", end.Sub(start))
	}

	{
		sum := 0
		start = time.Now()
		for i := 0; i < 10000000; i++ {
			sum += msi["a"]
		}
		end = time.Now()
		fmt.Println("map read time:", end.Sub(start))
	}

	{
		sum := 0
		start = time.Now()
		for i := 0; i < 10000000; i++ {
			sum += ssi.a
		}
		end = time.Now()
		fmt.Println("struct read time:", end.Sub(start))
	}

	// check write time
	{
		sum := 0
		current := 0
		start = time.Now()
		for i := 0; i < 10000000; i++ {
			sum += 1
			current = sum
		}
		end = time.Now()
		fmt.Println("int write time:", end.Sub(start), current)
	}

	{
		sum := 0
		start = time.Now()
		for i := 0; i < 10000000; i++ {
			sum += 1
			li[0] = sum
		}
		end = time.Now()
		fmt.Println("list write time:", end.Sub(start))
	}

	{
		sum := 0
		start = time.Now()
		for i := 0; i < 10000000; i++ {
			sum += 1
			msi["a"] = sum
		}
		end = time.Now()
		fmt.Println("map write time:", end.Sub(start))
	}

	{
		sum := 0
		start = time.Now()
		for i := 0; i < 10000000; i++ {
			sum += 1
			ssi.a = sum
		}
		end = time.Now()
		fmt.Println("struct write time:", end.Sub(start))
	}
}
