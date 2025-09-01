package main

import (
	"fmt"
)

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type List[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (l *List[T]) Add(t T) {
	n := &Node[T]{
		Value: t,
	}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}

	l.Tail.Next = n
	l.Tail = n
}

func (l *List[T]) Insert(t T, pos int) {
	n := &Node[T]{
		Value: t,
	}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}

	if pos <= 0 {
		n.Next = l.Head
		l.Head = n
		return
	}

	cn := l.Head
	for i := 1; i < pos; i++ {
		if cn.Next == nil {
			cn.Next = n
			l.Tail = n
			return
		}
		cn = cn.Next
	}

	if l.Tail == cn.Next {
		l.Tail = n
	}
	n.Next = cn.Next
	cn.Next = n
}

func (l *List[T]) Index(t T) int {
	i := 0
	for cn := l.Head; cn != nil; cn = cn.Next {
		if cn.Value == t {
			return i
		}
		i = i + 1
	}
	return -1
}

func (l *List[T]) Print() {
	for cn := l.Head; cn != nil; cn = cn.Next {
		fmt.Println(cn.Value)
	}
}

func main() {
	l := &List[int]{}
	l.Add(5)
	l.Add(10)
	l.Insert(8, 1)
	fmt.Println(l.Index(8))

	l.Print()
}
