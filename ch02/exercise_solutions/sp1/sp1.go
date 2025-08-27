package main

import "fmt"
import "time"

const cig = 1

var vig = 1

func main() {
	var start time.Time
	var end time.Time
	var sum int

	const cil = 1
	var vil = 1

	// time elapsed
	start = time.Now()
	end = time.Now()
	fmt.Println("time:", end.Sub(start))

	sum = 0
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		sum += 1
	}
	end = time.Now()
	fmt.Println("normal time:", end.Sub(start))

	sum = 0
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		sum += cig
	}
	end = time.Now()
	fmt.Println("cig time:", end.Sub(start))

	sum = 0
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		sum += vig
	}
	end = time.Now()
	fmt.Println("vig time:", end.Sub(start))

	sum = 0
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		sum += cil
	}
	end = time.Now()
	fmt.Println("cil time:", end.Sub(start))

	sum = 0
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		sum += vil
	}
	end = time.Now()
	fmt.Println("vil time:", end.Sub(start))

	// vig, vil に足し算をする
	vil = 0
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		vil++
	}
	end = time.Now()
	fmt.Println("vil add time:", end.Sub(start))

	vig = 0
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		vig++
	}
	end = time.Now()
	fmt.Println("vig add time:", end.Sub(start))

}
