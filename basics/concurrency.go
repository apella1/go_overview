package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	// ! spinning up a coroutine using the `go` keyword
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: %s\n", message)
	}()
	fmt.Printf("Email sent: %s\n", message)
}
