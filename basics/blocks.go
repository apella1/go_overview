package main

import "fmt"

func blocks() {
	// each place where a declaration occurs is a block
	// ! go places the predeclared identifiers in a universe block
	// ! the identifiers in the universe block can be shadowed in other scopes

	fmt.Println(true)
	// shadowing true
	true := "True text"
	fmt.Println(true)

	// ! universe block identifiers should never be redefined
}
