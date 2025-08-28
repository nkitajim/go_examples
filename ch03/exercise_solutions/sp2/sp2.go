package main

import (
	"fmt"
	"maps"
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
	var lio = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var msio = map[string]int{
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
	var ssio = Ssi{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var start time.Time
	var end time.Time

	{
		fmt.Println("-----------")

		li := make([]int, len(lio))
		msi := make(map[string]int, len(msio))
		ssi := Ssi{}
		copy(li, lio)
		maps.Copy(msi, msio)
		ssi = ssio

		fmt.Println("li:", li)
		fmt.Println("msi:", msi)
		fmt.Println("ssi:", ssi)
		fmt.Println()

		li1 := li
		msi1 := msi
		ssi1 := ssi

		li1[0] = 999
		msi1["a"] = 999
		ssi1.a = 999

		fmt.Println("li:", li)
		fmt.Println("msi:", msi)
		fmt.Println("ssi:", ssi)
		fmt.Println("li1:", li1)
		fmt.Println("msi1:", msi1)
		fmt.Println("ssi1:", ssi1)

	}

	{
		fmt.Println("-----------")
		li := lio
		msi := msio
		ssi := ssio

		start = time.Now()
		for i := 0; i < 10000000; i++ {
			li = lio
		}
		end = time.Now()
		li[0] = 1
		fmt.Println("shallow copy list time:", end.Sub(start))

		start = time.Now()
		for i := 0; i < 10000000; i++ {
			msi = msio
		}
		end = time.Now()
		msi["a"] = 1
		fmt.Println("shallow copy map time:", end.Sub(start))

		start = time.Now()
		for i := 0; i < 10000000; i++ {
			ssi = ssio
		}
		end = time.Now()
		ssi.a = 1
		fmt.Println("copy struct time:", end.Sub(start))
	}

	{
		fmt.Println("-----------")
		li := make([]int, len(lio))
		msi := make(map[string]int, len(msio))
		ssi := Ssi{}

		start = time.Now()
		for i := 0; i < 10000000; i++ {
			copy(li, lio)
		}
		end = time.Now()
		li[0] = 1
		fmt.Println("deep copy list time:", end.Sub(start))

		start = time.Now()
		for i := 0; i < 10000000; i++ {
			maps.Copy(msi, msio)
		}
		end = time.Now()
		msi["a"] = 1
		fmt.Println("deep copy map time:", end.Sub(start))

		start = time.Now()
		for i := 0; i < 10000000; i++ {
			ssi = ssio
		}
		end = time.Now()
		ssi.a = 1
		fmt.Println("copy struct time:", end.Sub(start))
	}
}
