package main

import "fmt"

// GitCommitHash is the SHA-1 hash of the git commit.
// The variable will be set by the Go linker tool. See
// the Makefile of this project.
var GitCommitHash string

func main() {
	fmt.Println("This binary was built from the following git commit hash:", GitCommitHash)
}
