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

func removeProfanity(message *string) {
	if message == nil {
		return
	}
	msgVal := *message
	msgVal = strings.ReplaceAll(msgVal, "dang", "****")
	msgVal = strings.ReplaceAll(msgVal, "shoot", "*****")
	msgVal = strings.ReplaceAll(msgVal, "heck", "****")
	*message = msgVal
}

// ! if a pointer points to nothing its zero value is nil

// * pointer receiver

type car struct {
	color string
}

func (c *car) setColor(color string) {
	c.color = color
}

type newEmail struct {
	message  string
	fromAddr string
	toAddr   string
}

func (e *newEmail) setMessage(newMsg string) {
	e.message = newMsg
}
