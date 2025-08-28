package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message := "Hi 👩 and 👨"
	runes := []rune(message)

	fmt.Printf("message len: %d\n", len(message))
	fmt.Printf("runes len: %d %d\n", len(runes), utf8.RuneCountInString(message))
	fmt.Println(message[3])
	fmt.Printf("%d %d %s\n", message[3], message[4], string(message[3]))

	fmt.Println(runes[3])
	fmt.Println(string(runes[3]))
}
