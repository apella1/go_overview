package main

import (
	"fmt"
)

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	// all the blocks of the loop are optional
	for i := 0; i < numMessages; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
	}
	return totalCost
}

func maxMessages(thresh float64) int {
	totalCost := 0.0
	// omitting the condition
	for i := 0; ; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
		if totalCost > thresh {
			return i
		}
	}
}

func getMaxMsgsToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 0
	for actualCostInPennies <= float64(maxCostInPennies) {
		maxMessagesToSend++
		actualCostInPennies *= costMultiplier
	}
	return maxMessagesToSend
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		word := ""
		if i%3 == 0 {
			word += "fizz"
		}
		if i%5 == 0 {
			word += "buzz"
		}
		if word != "" {
			fmt.Println(word)
		}
		fmt.Println(i)
	}
}

func loops() {
	// there is no while loop in go
	plantHeight := 1
	for plantHeight < 5 {
		fmt.Println("Plant height is", plantHeight, "Still growing.")
		plantHeight++
	}
}

func printPrimes(max int) {
	for i := 2; i <= max; i++ {
		if i == 2 {
			fmt.Println(i)
			continue
		}
		if i%2 == 0 {
			continue
		}
		isPrime := true
		for j := 3; j*j < i+1; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if !isPrime {
			continue
		}
		fmt.Println(i)
	}
}
