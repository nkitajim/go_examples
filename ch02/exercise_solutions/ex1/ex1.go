package main

import "fmt"

func main() {
	var i int = 20
	// var f float64 = i // vet error
	var f32 float32 = float32(i)
	var f64 float64 = float64(i)
	// fmt.Printf("i = %d, %f, f = %d, %f\n", i, i, f, f) // vet error. int is %f not support
	fmt.Printf("i = %d\n", i)
	fmt.Printf("f32 = %f\n", f32)
	fmt.Printf("f64 = %f\n", f64)
	fmt.Println(i)
	fmt.Println(f32)
	fmt.Println(f64)

	/* vet error, int is not float32
	if i == f32 {
		fmt.Println("OK")
	}
	*/
}
