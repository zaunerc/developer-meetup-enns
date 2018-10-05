package main

import (
	"fmt"
	"log"
	"os"
)

// Execute "$ node wasm_exec.js main.wasm" to run
// this code inside node.
func main() {
	fmt.Println("Hello world from Go again!")

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}
