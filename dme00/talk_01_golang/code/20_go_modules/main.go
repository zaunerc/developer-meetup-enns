package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// The dependency github.com/shopspring/decimal will be
// automatically fetched when building the program. The important
// thing here is to remember that you place the source code
// outside of your GOPATH. Otherwise "go build" won't use
// the new Go modules to manage your dependencies. If you want
// This might change at some later point. To vendor the depdency
// execute "go mod vendor".
func main() {
	dec, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", dec)
}
