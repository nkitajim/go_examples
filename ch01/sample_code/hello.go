package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Hello", "World!")
	// fmt.Println("Hello %s!", "World") // vet error, %sを使うべき関数ではない
	fmt.Println("")
	fmt.Printf("Hello %s!", "World")
	fmt.Printf("Hello %s!\n", "World")
	// fmt.printf("Hello %s!", "World") // not found
}
