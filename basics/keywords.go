package main

import "fmt"

func keyWords() {
	goKeywords := make([]string, 25)
	goKeywords = append(goKeywords, "const")
	goKeywords = append(goKeywords, "func")
	goKeywords = append(goKeywords, "struct")
	goKeywords = append(goKeywords, "interface")
	fmt.Println(goKeywords)
}
