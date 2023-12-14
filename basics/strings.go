package main

import "fmt"

func stringFormatting() {
	// formatting strings
	const name = "Theodore Roosevelt"
	const openRate = 46.2
	// Sprintf returns the string rather than outputting it to the stdout
	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)
	fmt.Println(msg)

	fmt.Printf("I am %v years old.", 22)
	//interpolating a string
	fmt.Printf("The %v is an amazing creature.", "low-flying eagle")
	// interpolating a decimal
	fmt.Printf("I am %f years old", 22.45)
}

// struct methods and string formatting
type authenticationInfo struct {
	username string
	password string
}

func (authInfo authenticationInfo) getBasicAuthInfo() string {
	return fmt.Sprintf("Authorization Basic %s:%s", authInfo.username, authInfo.password)
}
