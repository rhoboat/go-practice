package main

import "fmt"

// Generic functions define types directly after the function name
func SlicesIndex[S ~[]E, E comparable](collection S, element E) int {
	for i, v := range collection {
		if v == element {
			return i
		}
	}

	return -1
}

// Clever way to return a default value for any type
func getZero[T any] () T {
	var z T
	return z
}

// Generic types are structs
type cons[T any] struct {
	car T
	cdr *cons[T]
}

type List[T any] struct {
	head, tail *cons[T]
}

// Define methods on generic types
func (lst *List[T]) Push (v T) {
	// only 1 element in the list? wtf is this?
	if lst.tail == nil {
		lst.head = &cons[T]{car: v}
		lst.tail = lst.head // why??
	} else {
		lst.tail.cdr = &cons[T]{car: v}
		lst.tail = lst.tail.cdr
	}
}

func (lst *List[T]) Unshift () T {
	if lst.head != nil {
		el := lst.head.car
		lst.head = lst.head.cdr
		return el
	}
	return getZero[T]()
}

func (lst *List[T]) AllElements () []T {
	els := []T{}
	for e := lst.head; e != nil; e = e.cdr {
		els = append(els, e.car)
	}
	return els
}

func main() {
	var s = []string{"foo", "bar", "zoo"}
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())

	lst.Unshift()
	fmt.Println("list after unshift:", lst.AllElements())
}
