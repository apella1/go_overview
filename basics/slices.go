package main

import (
	"errors"
	"fmt"
)

func slices() {
	a := [7]int{1, 2, 3, 4, 5, 6, 7}
	// * slices are built on top of arrays
	b := a[2 : len(a)-1]
	fmt.Println(b)

	// creating a new slice using the make func
	// * func make([]T, len, cap) []T
	mySlice := make([]int, 5, 10)
	fmt.Println(mySlice)
	// capacity is usually omitted and defaults to the len
	otherSlice := make([]string, 6)
	fmt.Println(otherSlice)

	// * slice literals
	affirmation := []string{"I", "Love", "Go"}
	fmt.Println(affirmation)
	// builtin cap() and len() for slices
	// capacity is the max len of slice before reallocation is necessary
	fmt.Println(cap(affirmation), len(affirmation))
}

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()

	// ! slices reference arrays
	if plan == planPro {
		return allMessages[:], nil
	}
	if plan == planFree {
		return allMessages[0:2], nil
	}
	// nil acts as the zero-value for slices
	return nil, errors.New("unsupported plan")
}

func getMessageCosts(messages []string) []float64 {
	messageCosts := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		messageCosts[i] = float64(len(messages[i])) * 0.01
	}
	return messageCosts
}
