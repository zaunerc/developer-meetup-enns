// In Go it's idiomatic to communicate errors via an
// explicit, separate return value. This contrasts with
// the exceptions used in languages like Java and Ruby and
// the overloaded single result / error value sometimes
// used in C. Go's approach makes it easy to see which
// functions return errors and to handle them using the
// same language constructs employed for any other,
// non-error tasks.

package main

import "errors"
import "fmt"

// By convention, errors are the last return value and
// have type `error`, a built-in interface.
func divide(numerator float64, denominator float64) (float64, error) {

	if denominator == 0 {
		// `errors.New` constructs a basic `error` value
		// with the given error message.
		return -1, errors.New("denominator is 0. divsion by zero is not allowed")
	}

	// A `nil` value in the error position indicates that
	// there was no error.
	return numerator / denominator, nil
}

func main() {

	// Note that the use of an
	// inline error check on the `if` line is a common
	// idiom in Go code.
	if result, err := divide(1, 0); err != nil {
		fmt.Println("division failed:", err)
	} else {
		fmt.Println("division worked. result:", result)
	}

}
