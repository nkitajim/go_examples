package main

import "fmt"

// const value float32 = 20.1 // vet error
const value1 int = 20
const value2 float64 = 20

func main() {
	i := value1
	f := float64(value1)

	fmt.Println(i)
	fmt.Println(f)

	i = int(value2)
	f = value2

	fmt.Println(i)
	fmt.Println(f)
}
