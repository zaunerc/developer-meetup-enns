package main

import "fmt"

// It's possible to use custom types as `error`s by
// implementing the `Error()` method on them. Here's a
// variant on the example above that uses a custom type
// to explicitly represent an argument error.
type divByZeroError struct {
	numerator   float64
	denominator float64
}

func (e *divByZeroError) Error() string {
	return fmt.Sprintf("can not calculate %f / %f because divsion by 0 is not allowed", e.numerator, e.denominator)
}

func divide(numerator float64, denominator float64) (float64, error) {

	if denominator == 0 {
		// In this case we use `&divByZeroError` syntax to build
		// a new struct, supplying values for the two
		// fields `numerator` and `denominator`.
		return -1, &divByZeroError{numerator, denominator}
	}

	return numerator / denominator, nil
}

func main() {

	fmt.Printf("\n*** EXECUTING divide(10, 0) ***\n\n")

	if result, err := divide(10, 0); err != nil {
		fmt.Println("division failed:", err)
	} else {
		fmt.Println("division worked. result:", result)
	}

	fmt.Printf("\n*** EXECUTING divide(20, 0) ***\n\n")

	// If you want to programmatically use the data in
	// a custom error, you'll need to get the error as an
	// instance of the custom error type via type
	// assertion.
	_, err := divide(20, 0)
	if divByZeroErrorObj, ok := err.(*divByZeroError); ok {
		fmt.Printf("Divsion by 0 not possible. Maybe you want to try it the other way around :-) %f / %f\n",
			divByZeroErrorObj.denominator, divByZeroErrorObj.numerator)
	}

	fmt.Println()
}
