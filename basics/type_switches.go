package main

import "fmt"

// a type switch makes it easy to do several type assertions in a series
func typeSwitches() {
	printNumericValue(1)
	printNumericValue("1")
	printNumericValue(struct{}{})
}

func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}
