package main

import (
	"fmt"
	"time"
)

// select is similar to switch statements

func logMessages(chEmails, chSms chan string) {
	for {
		select {
		case email, ok := <-chEmails:
			if !ok {
				return
			}
			logEmail(email)
		case sms, ok := <-chSms:
			if !ok {
				return
			}
			logSms(sms)
		}
	}
}

func logSms(sms string) {
	fmt.Println("SMS:", sms)
}

func logEmail(email string) {
	fmt.Println("Email:", email)
}

// read-only channels
func saveBackups(snapshotTicker, saveAfter <-chan time.Time) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot()
		case <-saveAfter:
			saveSnapshot()
			return
		default:
			waitForData()
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func takeSnapshot() {
	fmt.Println("Taking Snapshot...")
}

func saveSnapshot() {
	fmt.Println("Saving Snapshot...")
}

func waitForData() {
	fmt.Println("Waiting for data")
}

// ! a send to nil channel blocks forever
// ! a receiver from a nil channel blocks forever
// zero value for a channel is nil
// ! a send to a closed channel panics
// ! a receive from a closed channel return the zero value immediately
