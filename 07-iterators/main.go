package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"
)

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

// Use iterators: https://pkg.go.dev/iter#Seq
func (lst *List[T]) AllElements () iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.cdr { // define how to iterate the structure
			if !yield(e.car) { // define when to stop
				return // break or early return to stop iterating
			}
		}
	}
}

// Define iterator without underlying data structure
func genFib() iter.Seq[int]{
	return func(yield func(int) bool) {
		for i, j := 0, 1; i <= 34; i, j = i + j, i {
			if !yield(i) { // so, you have to call yield on every iteration
				return
			}
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// call iterators in a for loop
	for v := range lst.AllElements() {
		fmt.Print(v, " ")
	}

	// slices.Collect takes an iter and returns a slice
	s := slices.Collect(lst.AllElements())
	fmt.Println(s)

	for f := range genFib() {
		fmt.Print(f, " ")
	}

	fmt.Println()
  // This is so cool. You don't need to build a slice first:
	for part := range strings.SplitSeq("go-by-example", "-") {
		fmt.Printf("part: %s\n", part)
	}
}
