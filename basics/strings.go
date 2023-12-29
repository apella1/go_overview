package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

func stringManipulation() {
	// * declaring multiline strings using backticks
	str := `This is
	a multiline
	string
	`
	fmt.Println(str)

	// * string concatenation
	// works okay
	newStr := "abc"
	newStr = newStr + "def"
	fmt.Print(newStr)

	// more efficient when combining lots of stings
	var sb strings.Builder
	sb.WriteString("This has been ")
	sb.WriteString("concatenated using strings Builder")
	fmt.Println(sb)

	// ! you can't simply cast a number to a string
	number := 124
	// castStr := string(number)
	// fmt.Println(castStr)
	rightStr := strconv.Itoa(number)
	fmt.Println(rightStr)
	otherRightStr := fmt.Sprintf("%d", number)
	fmt.Println(otherRightStr)

	// * strings can be converted into byte slices and byte slices can be converted to strings
	original := "this is a string"
	var b []byte
	b = []byte(original)
	fmt.Println(b)
	var s string
	s = string(b)
	fmt.Println(s)
}
