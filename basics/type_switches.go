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

func switchOverview() {
loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 2, 4, 6, 8:
			fmt.Println(i, "is even.")
		case 3:
			fmt.Println(i, "is divisible by 3 but not 2.")
		case 7:
			fmt.Println("exit the loop.")
			// breaking out of the outer for loop
			break loop
		default:
			fmt.Println(i, "is boring.")
		}
	}

	words := []string{"gopher", "many", "organizations", "decipher", "jettison"}
	for _, word := range words {
		// ! just like an if, variables can be scoped to the switch statement
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word.")
		case 5:
			// * block scoped within the case
			wordLen := len(word)
			fmt.Println(word, "is of the right length", wordLen)
		// nothing happens in case of an empty case
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is too long.")
		}
	}

	// * The blank switch
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 3:
			fmt.Println(word, "is a short word.")
		case wordLen > 3 && wordLen < 8:
			fmt.Println(word, "is of the right size.")
		case wordLen >= 8:
			fmt.Println(word, "is a long word.")
		}
	}
}

// ! using a blank switch to evaluate the expressions
func fizzBuzzSwitch() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
