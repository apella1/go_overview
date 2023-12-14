package main

import "fmt"

func anonymousFunctions() {
	primeNumbers := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(doMath(double, primeNumbers))

	// * anonymous funcs -> similar to js' arrow funcs
	tripleArrays := doMath(func(x int) int { return x * 3 }, primeNumbers)
	fmt.Println(tripleArrays)
}

func doMath(f func(x int) int, nums []int) []int {
	results := []int{}

	for _, num := range nums {
		results = append(results, f(num))
	}
	return results
}

func double(x int) int {
	return x * 2
}
