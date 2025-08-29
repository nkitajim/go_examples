package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	var start, end time.Time

	start = time.Now()
	p1 := make([]Person, 10000000)
	end = time.Now()
	fmt.Println("make p1 time", end.Sub(start), len(p1))

	start = time.Now()
	p2 := make([]Person, 0, 10000000)
	end = time.Now()
	fmt.Println("make p2 time", end.Sub(start), len(p2))

	start = time.Now()
	p3 := []Person{}
	for i := 0; i < 10000000; i++ {
		p3 = append(p3, Person{})
	}
	end = time.Now()
	fmt.Println("make p3 time", end.Sub(start), len(p3))

	start = time.Now()
	p4 := []Person{}
	for i := 0; i < 10000000; i++ {
		p4 = append(p4, Person{"Noboru", "Kitajima", 43})
	}
	end = time.Now()
	fmt.Println("make p4 time", end.Sub(start), len(p4))

	start = time.Now()
	p5 := make([]Person, 0, 10000000)
	for i := 0; i < 10000000; i++ {
		p5 = append(p5, Person{"Noboru", "Kitajima", 43})
	}
	end = time.Now()
	fmt.Println("make p5 time", end.Sub(start), len(p5))
}
