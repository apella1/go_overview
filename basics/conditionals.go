package main

import "fmt"

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
	if length := 56; length > 100 {
		fmt.Println("You have met the length requirements for the course!")
	}
}
