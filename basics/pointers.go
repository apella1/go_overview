package main

import (
	"fmt"
	"strings"
)

func pointers() {
	// * strPointer is initialized. *string is its type
	var strPointer *string
	myString := "hello"
	strPointer = &myString
	myStrPtr := &myString
	fmt.Println(strPointer)
	fmt.Printf("The type of myStrPtr is %T\n", myStrPtr) // *string
	// ! dereferencing the pointer to access underlying value
	fmt.Println(*myStrPtr)
	sendMessage(Message{
		Recipient: "Peter",
		Text:      "Good morning Peter!",
	})
}

/**
* passing values to functions to directly mutate them
* performance
* ! pointers are dangerous
 */

type Message struct {
	Recipient, Text string
}

func sendMessage(m Message) {
	fmt.Printf("To: %v\n", m.Recipient)
	fmt.Printf("Message: %v\n", m.Text)
}

func makeString(n int) string {
	word := ""
	for i := 0; i < n; i++ {
		word += "*"
	}
	return word
}

func removeProfanity(message *string) {
	profanityWords := []string{"dang", "shoot", "heck", "bitch"}
	profanityWords = append(profanityWords, "shit")
	if message == nil {
		return
	}

	msgVal := *message
	for _, p := range profanityWords {
		msgVal = strings.ReplaceAll(msgVal, p, makeString(len(p)))
	}
	*message = msgVal
}

// ! if a pointer points to nothing its zero value is nil

// * pointer receiver

type car struct {
	color      string
	noOfWheels int
}

func (c *car) setColor(color string) {
	c.color = color
}

func (c *car) setWheels(wheels int) {
	c.noOfWheels = wheels
}

func (c car) calculateSpeed() int {
	return c.noOfWheels * 40
}

type newEmail struct {
	message  string
	fromAddr string
	toAddr   string
}

func (e *newEmail) setMessage(newMsg string) {
	e.message = newMsg
}
