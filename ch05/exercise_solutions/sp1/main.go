package main

import (
	"fmt"
)

func bubbleSort(li []int) []int {
loop:
	loopCheck := false
	for ix := 0; ix < len(li)-1; ix++ {
		if li[ix] > li[ix+1] {
			tmp := li[ix]
			li[ix] = li[ix+1]
			li[ix+1] = tmp
			loopCheck = true
		}
	}
	if loopCheck {
		goto loop
	}
	return li
}

func bubbleSort2(li []int) []int {
	for {
		loopCheck := true
		for ix := 0; ix < len(li)-1; ix++ {
			if li[ix] > li[ix+1] {
				tmp := li[ix]
				li[ix] = li[ix+1]
				li[ix+1] = tmp
				loopCheck = false
			}
		}
		if loopCheck {
			break
		}
	}
	return li
}

func bubbleSortEx(li []int, check func(int, int) bool) []int {
loop:
	loopCheck := false
	for ix := 0; ix < len(li)-1; ix++ {
		if check(li[ix], li[ix+1]) {
			tmp := li[ix]
			li[ix] = li[ix+1]
			li[ix+1] = tmp
			loopCheck = true
		}
	}
	if loopCheck {
		goto loop
	}
	return li
}

func main() {
	li := []int{3, 2, 5, 6, 7, 8, 1}
	li2 := []int{3, 2, 5, 6, 7, 8, 1}
	li3 := []int{3, 2, 5, 6, 7, 8, 1}
	li4 := []int{3, 2, 5, 6, 7, 8, 1}

	sli := bubbleSort(li)
	sli2 := bubbleSort2(li2)
	sli3 := bubbleSortEx(li3, func(ix int, iy int) bool {
		if ix > iy {
			return true
		}
		return false
	})
	sli4 := bubbleSortEx(li4, func(ix int, iy int) bool {
		if ix < iy {
			return true
		}
		return false
	})
	fmt.Println(li)
	fmt.Println(sli)
	fmt.Println(sli2)
	fmt.Println(sli3)
	fmt.Println(sli4)
}
