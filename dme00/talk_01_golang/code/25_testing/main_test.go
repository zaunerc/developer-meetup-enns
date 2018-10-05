package main

import "testing"

// Go unit tests do not require any specific frameworks or syntax.
func TestSum(t *testing.T) {

	testTables := []struct {
		inputA   int
		inputB   int
		expected int
	}{
		{
			1, 2, 3,
		},
		{
			2, 2, 4,
		},
		{
			3, 3, 6,
		},
	}

	for _, testTable := range testTables {
		actualResult := Sum(testTable.inputA, testTable.inputB)
		if actualResult != testTable.expected {
			t.Errorf("\nExpected:\t%v\nGot:\t\t%v", testTable.expected, actualResult)
		}
	}
}
