package main

import "fmt"

func strings() {
	// formatting strings
	const name = "Theodore Roosevelt"
	const openRate = 46.2
	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)
	fmt.Println(msg)

	fmt.Printf("I am %v years old.", 22)
	//interpolating a string
	fmt.Printf("The %v is an amazing creature.", "low-flying eagle")
	// interpolating a decimal
	fmt.Printf("I am %f years old", 22.45)
}
