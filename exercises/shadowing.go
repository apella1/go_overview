package exercises

import "fmt"

func Shadowing() {
	// * a shadowing variable is a variable that has the same name
	// * as a variable in a containing block
	x := 10
	if x > 5 {
		fmt.Println(x)
		// ! this shadows the outer x within the if block
		// ! shadowing should only be when intended
		x, y := 5, 20
		fmt.Println(x, y)
	}
	fmt.Println(x)
}
