package main

import "fmt"

func arraySum(nums ...int) int {
	// ! ...int == []int -> the difference is on the caller being able to pass
	// ! any number of args
	// fmt.Println() and fmt.Sprintf() are variadic
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func variadicFuncs() {
	names := []string{"Peter", "James", "John", "Janice"}
	// using spread operator to pass in the args
	printStrings(names...)

	girlNames := []string{"Mary", "Angie", "Dorcas"}
	boyNames := []string{"Kelly", "Lane", "John"}
	// append() is a variadic function for resizing slices
	allNames := append(girlNames, boyNames...)
	fmt.Println(allNames)
}

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	dailyCosts := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		if cost.day >= len(dailyCosts) {
			dailyCosts = append(dailyCosts, 0.0)
		}
		dailyCosts[cost.day] += cost.value
	}
	return dailyCosts
}
