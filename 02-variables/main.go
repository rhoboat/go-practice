package main

import "fmt"

var outside string

func main() {

	// uninitialized types
	var boolean bool
	fmt.Println("boolean=", boolean)

	var float float32
	fmt.Println("float=", float)

	var a, b, c int = 1, 2, 3
	fmt.Println("a=", a, "b=", b, "c=", c)

	d, e, f := "d", "e", "f"
	fmt.Println("d=", d, "e=", e, "f=", f)

	var emptyString string
	fmt.Println("emptyString=", emptyString)

	fmt.Println("outside=", outside)
	outside = "set inside"
	fmt.Println("outside=", outside)
}
