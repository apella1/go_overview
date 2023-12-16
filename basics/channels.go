package main

import (
	"fmt"
	"time"
)

// channels are go-routine safe queues
// ! deadlocks are when all the go routines in the program are blocking

/* type otherEmail struct {
	Date Date
}

func filterOldEmails(emails []otherEmail) {
	isOldChan := make(chan bool)

	// spinning up a new go routine
	go func() {
		for _, e := range emails {
			if e.Date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
				isOldChan <- true
				continue
			}
			isOldChan <- false
		}
	}()

	isOld := <-isOldChan
	fmt.Println("email 1 is old:", isOld)
	isOld = <-isOldChan
	fmt.Println("email 2 is old:", isOld)
	isOld = <-isOldChan
	fmt.Println("email 3 is old:", isOld)
}
*/

// tokens are unary values
// empty structs are tokens
func waitForDbs(numDbs int, dbChan chan struct{}) {
	for i := 0; i < numDbs; i++ {
		<-dbChan
	}

}

func getDatabasesChannel(numDbs int) chan struct{} {
	ch := make(chan struct{})
	go func() {
		for i := 0; i < numDbs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online.", i+1)
		}
	}()
	return ch
}

// * buffered channels
func addEmailsToQueue(emails []string) chan string {
	emailsToSend := make(chan string, len(emails))
	for _, email := range emails {
		emailsToSend <- email
	}
	return emailsToSend
}

func sendEmails(batchSize int, ch chan string) {
	for i := 0; i < batchSize; i++ {
		email := <-ch
		fmt.Println("Sending email:", email)
	}
}

// ! channels in go can be closed
// - a channel should always be closed from the sending side
// ! don't send a value on a closed channel

func countReports(numSentChan chan int) int {
	total := 0
	// infinite for loop
	for {
		numSent, ok := <-numSentChan
		if !ok {
			break
		}
		total += numSent
	}
	return total
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
		fmt.Printf("Send batch of %v reports\n", numReports)
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}

// * range keyword works in channels as in slices and maps
func concurrentFib(n int) {
	chInts := make(chan int)
	go func() {
		fibonacci(n, chInts)
	}()
	for v := range chInts {
		fmt.Println(v)
	}

}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}
