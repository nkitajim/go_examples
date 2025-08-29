package main

import (
	"fmt"
)

type Logger interface {
	Log(s string)
}

type Log1 struct{}

func (l Log1) Log(s string) {
	fmt.Println("[Log1]", s)
}

type Log2 struct{}

func (l Log2) Log(s string) {
	fmt.Println("[Log2]", s)
}

func Log3(s string) {
	fmt.Println("[Log3]", s)
}

func Log4(s string) {
	fmt.Println("[Log4]", s)
}

func main() {
	{
		log := Log1{}
		log.Log("hello")
	}

	{
		log := Log2{}
		log.Log("hello")
	}

	{
		Log3("hello")
		Log4("hello")
	}
}
