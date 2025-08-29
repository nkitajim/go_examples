package main

import (
	"fmt"
)

func UpdateSlice(ss []string, s string) {
	ss[len(ss)-1] = s
	fmt.Println(ss)
}

func GrowSlice(ss []string, s string) {
	ss = append(ss, s)
	fmt.Println(ss)
}

func ReturnSlice(ss []string, s string) []string {
	ss = append(ss, s)
	fmt.Println(ss)
	return ss
}

func CapSlice(ss []string, s string) {
	ss = append(ss, s)
	fmt.Println(ss)
}

func main() {
	us := []string{"Hello", "World", "!"}
	fmt.Println("org:", us)
	UpdateSlice(us, "end")
	fmt.Println("after:", us)

	gs := []string{"Hello", "World", "!"}
	fmt.Println("org:", gs)
	GrowSlice(gs, "end")
	fmt.Println("after:", gs)

	rs := []string{"Hello", "World", "!"}
	fmt.Println("org:", rs)
	rs = ReturnSlice(rs, "end")
	fmt.Println("after:", rs)
}
