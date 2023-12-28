package exercises

// * Exercises from `Learning Go` by Jon Bodner

import "fmt"

func PrintIntAndFloat() {
	i := 20
	f := 34.66
	i = int(f)
	fmt.Println(i, f)
}

func MultipleAssignments() {
	const value = 30
	var nonConstValue int = 23
	var i int = value
	var f float64 = float64(value)
	fmt.Println(i, f)
	fmt.Println(float64(nonConstValue))
}

func MultipleVariables() {
	// * byte is an alias for uint8
	// * operations can be done between uint8 and byte
	// * uint8 is however uncommon thus byte is used in its place
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	// *  adding 1 to the max value for all the types
	fmt.Println(b+1, smallI+1, bigI+1)
}

// todo bitwise operations and reassignments