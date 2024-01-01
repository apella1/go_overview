package main

type person struct {
	name string
	age  int
}

// maps and slices behave differently than the structs
// and other primitive types because they are implemented
// with pointers
func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}
