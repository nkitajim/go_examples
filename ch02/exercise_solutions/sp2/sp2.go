package main

import "fmt"
import "time"

func main() {
	var start time.Time
	var end time.Time

	// time elapsed
	start = time.Now()
	end = time.Now()
	fmt.Println("time:", end.Sub(start))

	var i16sum int16 = 0
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		i16sum += 1
	}
	end = time.Now()
	fmt.Println("i16sum time:", end.Sub(start))

	var i32sum int32 = 0
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		i32sum += 1
	}
	end = time.Now()
	fmt.Println("i32sum time:", end.Sub(start))

	var f32sum float32 = 0
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		f32sum += 1
	}
	end = time.Now()
	fmt.Println("f32sum time:", end.Sub(start))

	var f64sum float64 = 0
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		f64sum += 1
	}
	end = time.Now()
	fmt.Println("f64sum time:", end.Sub(start))

	var i16mul int16 = 0
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		i16mul *= 2
	}
	end = time.Now()
	fmt.Println("i16mul time:", end.Sub(start))

	var i32mul int32 = 0
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		i32mul *= 2
	}
	end = time.Now()
	fmt.Println("i32mul time:", end.Sub(start))

	var f32mul float32 = 0
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		f32mul *= 2
	}
	end = time.Now()
	fmt.Println("f32mul time:", end.Sub(start))

	var f64mul float64 = 0
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		f64mul *= 2
	}
	end = time.Now()
	fmt.Println("f64mul time:", end.Sub(start))
}
