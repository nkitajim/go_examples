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
		var i16 int16 = 1
		i16sum += i16
	}
	end = time.Now()
	fmt.Println("i16sum time:", end.Sub(start))

	var i32sum int32 = 0
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		var i32 int32 = 1
		i32sum += i32
	}
	end = time.Now()
	fmt.Println("i32sum time:", end.Sub(start))

	var f32sum float32 = 0
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		var f32 float32 = 1
		f32sum += f32
	}
	end = time.Now()
	fmt.Println("f32sum time:", end.Sub(start))

	var f64sum float64 = 0
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		var f64 float64 = 1
		f64sum += f64
	}
	end = time.Now()
	fmt.Println("f64sum time:", end.Sub(start))
}
