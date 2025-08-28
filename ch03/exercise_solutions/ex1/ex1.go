package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "こんばんは", "こんにち", "おはよう"}
	// var greetings []string = []string{"Hello", "Hola", "こんばんは", "こんにち", "おはよう"}
	// var greetings = []string{"Hello", "Hola", "こんばんは", "こんにち", "おはよう"}
	slice1 := greetings[:2]
	slice2 := greetings[1:4]
	slice3 := greetings[3:]

	fmt.Println(greetings)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
