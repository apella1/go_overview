package main

import "fmt"

func variables() {
	// initializing variables
	// the standard sizes should be used unless there is a need for the use case
	// i.e int uint float64 complex128

	var smsSendingLimit uint64
	var costPerSms float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSms, hasPermission, username)

	accountAge := 2.6
	// type casting from float to int is truncated
	accountAgeInt := int(accountAge)
	fmt.Println("Your account has existed for", accountAgeInt, "years.")

	// constants do not accept the shorthand syntax
	// not similar to js
	const premiumPlanName = "Top Achievers"
	const basicPlanName = "Starters"

	// computed constants
	const firstName = "Mike"
	const lastName = "Otis"
	const fullName = firstName + " " + lastName

	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour
	fmt.Println("The number of seconds in an hours is", secondsInHour)

	// variables are passed by value, not by reference
	x := 7
	increment(x)
	fmt.Println(x)
}

func increment(a int) {
	a++
}
