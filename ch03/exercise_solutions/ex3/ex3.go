package main

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	id        int
}

type employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	e1 := Employee{
		"Noboru",
		"Kitajima",
		1,
	}
	e2 := Employee{
		firstName: "Noboru",
		lastName:  "Kitajima",
		id:        1,
	}
	var e3 Employee
	e3.firstName = "Noboru"
	e3.lastName = "Kitajima"
	e3.id = 1

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)

	fmt.Println(employee{"Noboru", "Kitajima", 1})
}
