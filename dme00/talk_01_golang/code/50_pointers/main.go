package main

import (
	"fmt"
)

func main() {

	fmt.Printf("\n*** changeToAbsoluteByVal ***\n\n")

	a := -10
	fmt.Printf("a before changeToAbsoluteByVal: %v\n", a)
	fmt.Printf("memory address of a: %p\n", &a)
	changeToAbsoluteByVal(a)
	fmt.Printf("a after changeToAbsoluteByVal: %v\n", a)

	fmt.Printf("\n*** changeToAbsoluteByPtr ***\n\n")

	b := -20
	fmt.Printf("b before changeToAbsoluteByVal: %v\n", a)
	fmt.Printf("memory address of b: %p\n", &b)
	changeToAbsoluteByPtr(&b)
	fmt.Printf("b after changeToAbsoluteByPtr: %v\n", b)

	fmt.Println()
}

func changeToAbsoluteByVal(val int) {
	fmt.Printf("memory address of given param in function changeToAbsoluteByVal: %p\n", &val)
	if val < 0 {
		val = val * -1
	}
}

func changeToAbsoluteByPtr(val *int) {

	// The * operator denotes the pointer's underlying value.
	// This is known as "dereferencing" or "indirecting".
	// Unlike C, Go has no pointer arithmetic.
	// This really makes sense when working with structs.

	fmt.Printf("memory address of given param in function changeToAbsoluteByPtr: %p\n", &*val)
	if *val < 0 {
		*val = *val * -1
	}
}
