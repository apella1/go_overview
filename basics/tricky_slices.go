package main

import "fmt"

func trickySlices() {
	fmt.Println("Tricky Slices Start Here")
	a := make([]int, 3)
	fmt.Println("Length of of a is", len(a))
	fmt.Println("Cap of of a is", cap(a))
	b := append(a, 4)
	fmt.Println("b:", b)
	fmt.Println("Address of b:", &b[0])
	c := append(a, 5)
	fmt.Println("Address of c:", &c[0])
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	// using range to iterate over slices
	for index, value := range c {
		fmt.Println(index, value)
	}
}

func fishySlices() {
	// ! the tricky part is when the slice is not reallocated
	// ! then the values are mutated as the underlying array remains the same
	i := make([]int, 3, 7)
	fmt.Println("Length of of i is", len(i))
	fmt.Println("Cap of of i is", cap(i))
	j := append(i, 4)
	fmt.Println("j:", j)
	fmt.Println("Address of j:", &j[0])
	k := append(i, 5)
	fmt.Println("Address of k:", &k[0])
	fmt.Println("i:", i)
	fmt.Println("j:", j)
	fmt.Println("k:", k)

}

func indexOfFirstBadWord(msg []string, badWords []string) int {
	var ind int
	for index, badWord := range badWords {
		found := false
		for _, word := range msg {
			if badWord == word {
				ind = index
				found = true
				break
			}
		}
		if found {
			return ind
		}
	}
	return -1
}
