package main

import (
	"errors"
	"fmt"
)

func functions() {
	fmt.Println(addNumbers(4, 7))
	fmt.Println(concatStr("Hello", " world"))
}

func addNumbers(x int, y int) int {
	return x + y
}

func concatStr(s1 string, s2 string) string {
	return s1 + s2
}

// when multiple parameters are of the same type,
// the type can be specified just after the last one (syntactic sugar)

func constructVehicle(name, model, company string, horsePower, wheels int) {
	fmt.Printf("%v is of model %v and is manufactured by %v. It has horse power of %v and has %v wheels", name, model, company, horsePower, wheels)
}

// ignoring function return values
// x and y are named return values
func getPoints() (x, y int) {
	return 3, 5
}

// ignoring y value
// x, _ := getPoints()

func getCoords() (x, y int) {
	// x and y are initialized with their zero values
	return // automatically return x and y
	// ! the implicit return however affects readability
}

// getCoords func expanded would be
func getCoordsEx() (int, int) {
	var x int
	var y int
	return x, y
}

// named return types
func yearUntilEvents(age int) (yearsUntilAdult, yearsToDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - 0
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}

	yearsToDrinking = 21 - 0
	if yearsToDrinking < 0 {
		yearsToDrinking = 0
	}

	yearsUntilCarRental = 28 - 0
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}

	return yearsUntilAdult, yearsToDrinking, yearsUntilCarRental
}

func calculator(a, b int) (product, quotient int, err error) {
	// ! early returns aka guard clauses
	if b == 0 {
		return 0, 0, errors.New("Can't divide by zero")
	}

	product = a * b
	quotient = a / b

	return product, quotient, nil
}
