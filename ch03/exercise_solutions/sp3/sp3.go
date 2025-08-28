package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Employee struct {
	ID        int    `validate:"gt=0"`
	FirstName string `validate:"required,max=16"`
	LastName  string `validate:"required,max=16"`
}

type employee struct {
	id        int    `validate:"gt=0"`
	firstName string `validate:"required,max=16"`
	lastName  string `validate:"required,max=16"`
}

func validation[T any](e T) error {
	fmt.Println(e)
	validate := validator.New()
	if err := validate.Struct(e); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag(), err.Param())
		}
		return err
	} else {
		fmt.Println("Validate safe")
	}
	return nil
}

func main() {
	e1 := Employee{
		ID:        1,
		FirstName: "Noboru",
		LastName:  "Kitajima",
	}
	validation(e1)

	e2 := Employee{
		ID:        0,
		FirstName: "hogehogehogehogehogehoge",
		LastName:  "Kitajima",
	}
	validation(e2)

	e3 := Employee{
		ID:        0,
		FirstName: "Noboru",
		// LastName:  "Kitajima",
	}
	validation(e3)

	e4 := employee{
		id: 0,
	}
	validation(e4)
}
