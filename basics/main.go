package main

// package main signifies that the program can be run as a standalone executable

import (
	"fmt"

	"github.com/apella1/go_overview/exercises"
	"github.com/apella1/go_overview/intermediate"
)

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

	sendsSoFar := 67
	const sendsToAdd = 322
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("You have sent", sendsSoFar, "messages.")

	// getting just the first name
	firstName, _ := getNames()
	fmt.Println("Welcome,", firstName)

	// myStrings := printStrings("Hello", "world", "majestic")
	// fmt.Println(myStrings)
	variadicFuncs()
	tenByTenMatrix := createMatrix(10, 10)
	fmt.Println(tenByTenMatrix)
	fmt.Println(divide(4, 0))
	trickySlices()
	fishySlices()
	slices()

	fmt.Println()
	fmt.Println("Maps start here!")
	maps()
	nestedMaps()
	keyTypes()
	structs()

	names := []string{"john", "jacob", "paul", "peter", "john", "john"}
	fmt.Println(getNameCounts(names))

	fmt.Println()
	fmt.Println("Advanced Functions")
	advancedFunctions()
	fmt.Println()
	fmt.Println("Function Closures")
	closures()
	anonymousFunctions()
	fmt.Println()
	fmt.Println("Pointers")
	pointers()

	sendEmail("Good morning")

	// From intermediate package
	intermediate.IntermediateConcepts()

	concurrentFib(30)
	generics()

	fmt.Println()
	fmt.Println("String manipulation")
	stringFormatting()
	stringManipulation()

	fmt.Println()
	fmt.Println("Exercises Package")
	exercises.PrintIntAndFloat()
	exercises.MultipleAssignments()
	exercises.MultipleVariables()
	exercises.ComparingMaps()

}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar
}

func getNames() (string, string) {
	return "John", "Doe"
}
