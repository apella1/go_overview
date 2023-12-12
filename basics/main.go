package main

// package main signifies that the program can be run as a standalone executable

import "fmt"

func main() {
	fmt.Println("Starting TextIo server!")
	messagesFromDoris := []string{
		"Learn Go",
		"Finish up Node course",
		"Go to the Gym",
		"Finally, write yourself a lovely note",
	}
	numOfMessages := float64(len(messagesFromDoris))
	costPerMsg := .02
	totalCost := numOfMessages * costPerMsg
	fmt.Printf("The type of costPerMsg is %T\n", costPerMsg)
	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)

	// variable assignment
	var username string = "king"
	var password string = "3523405723"
	fmt.Println("Authorization: Basic", username+":"+password)

	// short variable declaration
	// Go infers the types of the variables
	play := "Hamlet"
	yearPlayPublished := 1989
	fmt.Println(play, "was published in the year", yearPlayPublished)

	// same line variable declaration
	language, author, isPublished := "English", "Dale Carnegie", false
	fmt.Println(language, author, isPublished)
}
