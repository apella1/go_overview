package main

import "fmt"

// * a closure is a func that references variables from outside its body

func closures() {
	wordMaker := concatter()
	wordMaker("Go")
	wordMaker("is")
	wordMaker("an")
	wordMaker("interesting")
	wordMaker("programming")
	fmt.Println(wordMaker("language."))

	myAdder := adder()
	myAdder(4)
	myAdder(5)
	myAdder(7)
	myAdder(9)
	fmt.Println(myAdder(12))
}

func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

type emailBill struct {
	costInPennies int
}

func adder() func(int) int {
	sum := 0
	return func(cost int) int {
		sum += cost
		return sum
	}
}
