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

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func main() {
	p1 := MakePerson("Noboru", "Kitajima", 43)
	p2 := MakePersonPointer("Kyoko", "Kitajima", 43)

	fmt.Println(p1)
	fmt.Println(*p2)

	var start, end time.Time
	var p string

	start = time.Now()
	for i := 0; i < 1000000; i++ {
		p = p1.FirstName
	}
	end = time.Now()
	fmt.Println(p, end.Sub(start))

	start = time.Now()
	for i := 0; i < 1000000; i++ {
		p = (*p2).FirstName
		// p = p2.FirstName // OK
	}
	end = time.Now()
	fmt.Println(p, end.Sub(start))
}
