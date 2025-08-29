package main

import (
	"sort"
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	for ix := 0; ix < b.N; ix++ {
		li := []int{3, 2, 5, 6, 7, 8, 1}
		bubbleSort(li)
	}
}

func BenchmarkBubbleSort2(b *testing.B) {
	for ix := 0; ix < b.N; ix++ {
		li := []int{3, 2, 5, 6, 7, 8, 1}
		bubbleSort2(li)
	}
}

func BenchmarkBubbleSortEx(b *testing.B) {
	for ix := 0; ix < b.N; ix++ {
		li := []int{3, 2, 5, 6, 7, 8, 1}
		bubbleSortEx(li, func(ix int, iy int) bool {
			if ix > iy {
				return true
			}
			return false
		})
	}
}

func BenchmarkSort(b *testing.B) {
	for ix := 0; ix < b.N; ix++ {
		li := []int{3, 2, 5, 6, 7, 8, 1}
		sort.Ints(li)
	}
}
