package main

import (
	"fmt"
)

//! error handling is built around the `error` interface in go
/* type error interface {
	Error() string
} */
// * error handling relies heavily on proper string formatting
// errors are nullable strings
// errors are not thrown

func errorHandling() {
	const name = "Abraham Lincoln"
	const age = 77
	s := fmt.Sprintf("%v is %v years old.", name, age)
	fmt.Println(s)
}

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	csmCost, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}

	spsCost, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0.0, err
	}

	return csmCost + spsCost, nil
}

func sendSMS(msg string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = .0002

	if len(msg) > maxTextLen {
		return 0.0, fmt.Errorf("Can't send messages with over %v characters.", maxTextLen)
	}
	return costPerChar * float64(len(msg)), nil
}

func getSMSErrStr(cost float64, recipient string) error {
	return fmt.Errorf("SMS that cost $%.4f cannot be sent to %v", cost, recipient)
}

// building custom error types
type userError struct {
	name string
}

func (e userError) Error() string {
	return fmt.Sprintf("%s has an error with their account", e.name)
}

// handling division by 0 error
type divErr struct {
	dividend string
}

func (e divErr) Error() string {
	return fmt.Sprintf("Cannot divide %v by zero!", e.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		// todo: how to represent the err method on the divErr struct
		return 0.0, divErr{}
	}
	return dividend / divisor, nil
}
