// _Interfaces_ are named collections of method
// signatures.

package main

import "fmt"
import "math"

// Here's a basic interface for geometric shapes.
type geometry interface {
	area() float64
	perim() float64
}

// For our example we'll implement this interface on
// `rect` and `circle` types.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// implement `geometry` on `rect`s.
//
// This is called a value receiver.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// The implementation for `circle`s.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type, then we can call
// methods that are in the named interface. Here's a
// generic `printMeasurementsToConsole` function taking advantage of this
// to work on any `geometry`.
func printMeasurementsToConsole(g geometry) {
	fmt.Printf("The area of the given geometry shape is %v\n", g.area())
	fmt.Printf("The perimeter of the given geometry shape is %v\n", g.perim())
}

func main() {
	rectangleObj := rect{width: 3, height: 4}
	circleObj := circle{radius: 5}

	// The `circle` and `rect` struct types both
	// implement the `geometry` interface so we can use
	// instances of
	// these structs as arguments to `measure`.

	fmt.Printf("\n*** rectangle ***\n\n")
	printMeasurementsToConsole(rectangleObj)
	fmt.Printf("\n*** circle ***\n\n")
	printMeasurementsToConsole(circleObj)

	fmt.Println()
}
