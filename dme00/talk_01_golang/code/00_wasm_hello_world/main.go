package main

import (
	"fmt"
)

// compile with $ GOOS=js GOARCH=wasm go build -o main.wasm main.go
func main() {
	fmt.Println("Hello world from Go!")
}
