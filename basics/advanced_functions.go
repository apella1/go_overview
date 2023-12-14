package main

import (
	"fmt"
	"io"
	"os"
)

// * functions as first-class citizens and higher order funcs
func advancedFunctions() {
	fmt.Println(aggregate(3, 5, 5, add))
	currying()
}

func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessages := []string{}
	for index, message := range messages {
		formattedMessages[index] = formatter(message)
		// *alt* formattedMessages = append(formattedMessages, formatter(message))
	}
	return formattedMessages
}

func addSignature(msg string) string {
	return fmt.Sprintf("%v. Kind regards", msg)
}

// ! currying
// * a function taking a function as input and returning another function

func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}

func currying() {
	squareNum := selfMath(multiply)
	doubleNum := selfMath(add)

	fmt.Println(doubleNum(2))
	fmt.Println(squareNum(9))
}

// ! the returned function is void
func getLogger(formatter func(string, string) string) func(string, string) {
	return func(x, y string) {
		fmt.Println(formatter(x, y))
	}
}

// ! defer keyword
func copyFile(dstName, srcName string) (written int64, err error) {
	// open source file
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	// closing the source file when the copyFile function returns
	defer src.Close()

	// create the destination file
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	// close the destination file when the copyFile func returns
	defer dst.Close()

	return io.Copy(dst, src)
}

const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

type userLevels struct {
	user, admin bool
}

func logAndDelete(users map[string]userLevels, name string) (log string) {
	defer delete(users, name)
	user, ok := users[name]
	if !ok {
		return logNotFound
	}
	if user.admin {
		return logAdmin
	}
	return logDeleted
}
