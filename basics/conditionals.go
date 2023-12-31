package main

import (
	"fmt"
	"math/rand"
)

func conditionals() {
	marks := 80
	passMark := 75
	fmt.Printf("You have scored %v marks for your paper.", marks)

	if marks >= passMark {
		fmt.Println("You have passed!")
	} else {
		fmt.Println("Sorry, you didn't make the cut. Try again next time.")
	}

	// else if
	if marks > 90 {
		fmt.Println("Your results are exquisite!")
	} else if marks >= passMark {
		fmt.Println("You have passed!")
	} else {
		fmt.Println("You have failed")
	}

	// syntactic sugar for limiting scope and keeping code concise
	// * limiting scope to the if statement
	if length := 56; length > 100 {
		fmt.Println("You have met the length requirements for the course!")
	}

	if n := rand.Intn(30); n == 0 {
		fmt.Println("That's too low")
	} else if n > 20 {
		fmt.Printf("%v is a very large number", n)
	} else {
		fmt.Printf("%v  is a good number.", n)
	}

	if choice := rand.Intn(3); choice == 1 {
		fmt.Println("You have selected", choice)
	} else if choice == 2 {
		fmt.Println(choice, "is even.")
	} else {
		fmt.Println(choice, "is odd.")
	}
}
