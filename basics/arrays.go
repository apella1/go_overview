package main

import "fmt"

func arrays() {
	// declaring an array of 10 integers
	var myInts [10]int
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(myInts)
	fmt.Println(primes)
}

func getMessageWithRetries() [3]string {
	messages := [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
	return messages
}

func send(name string, doneAt int) {
	fmt.Printf("Sending to %v", name)
	fmt.Println()

	messages := getMessageWithRetries()
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf("Sending %v", msg)
		fmt.Println()

		if i == doneAt {
			fmt.Println("They responded!")
			break
		}
		if i == len(messages)-1 {
			fmt.Println("Complete failure")
		}
	}
}
