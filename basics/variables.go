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
}
